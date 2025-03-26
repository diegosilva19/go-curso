package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"example.com/cli-project/services"
	"k8s.io/utils/env"
)

// RequestBody represents each command object posted to /add_commands.
type RequestBody struct {
	Command   string `json:"command"`
	WriteFile string `json:"write_file"`
}

// commandInfo holds everything we need to track and manage the command.
type commandInfo struct {
	ID        string    `json:"id"`
	Command   string    `json:"command"`
	WriteFile string    `json:"write_file"`
	IsRunning bool      `json:"is_running"`
	cmd       *exec.Cmd // Not exposed in JSON, used internally to kill the process.
}

var (
	// wg tracks all goroutines running commands, so we can wait on graceful shutdown if needed.
	wg sync.WaitGroup

	// commandsMap maps each command's unique ID -> commandInfo
	commandsMap = struct {
		sync.RWMutex
		data map[string]*commandInfo
	}{data: make(map[string]*commandInfo)}

	// atomic counter to help generate unique IDs
	counter int64
)

// generateID creates a unique ID from a timestamp + a counter.
func generateID() string {
	idCounter := atomic.AddInt64(&counter, 1)
	return fmt.Sprintf("%d-%d", time.Now().UnixNano(), idCounter)
}

// Add configuration struct
type Config struct {
	Port           string `json:"port"`
	WriteDirectory string `json:"write_directory"`
	MaxConcurrency int    `json:"max_concurrency"`
}

var cfg = loadConfig()

func registerEndpoints() {
	http.HandleFunc("/add_commands", addCommandsHandler)
	http.HandleFunc("/status_commands", statusCommandsHandler)
	http.HandleFunc("/delete_command", deleteCommandHandler)
	http.HandleFunc("/delete_all_commands", deleteAllCommandsHandler)
	http.HandleFunc("/kube/retrieve_list", services.CreateList)
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/ready", readinessCheck)
}

func main() {

	registerEndpoints()

	// Create server with timeouts
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Setup signal handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("Server started on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("Server error: %v\n", err)
		}

	}()
	// Wait for shutdown signal
	<-stop

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown error: %v\n", err)
	}

	// Wait for running commands to finish
	wg.Wait()
}

// Add configuration loading
func loadConfig() *Config {

	maxConcurrency, err := env.GetInt("MAX_CONCURRENCY", 100)
	if err != nil {
		// Fall back to default value if there's an error
		maxConcurrency = 100
	}

	defaultWriteDirm := env.GetString("WRITE_DIR", "/var/log/commands")

	if defaultWriteDirm == "" {
		ex, err := os.Executable()
		if err != nil {
			fmt.Printf("Failed to get executable path: %v\n", err)
			panic(err)
		}

		exPath := filepath.Dir(ex)
		defaultWriteDirm = exPath
	}

	return &Config{
		Port:           env.GetString("PORT", "8080"),
		WriteDirectory: defaultWriteDirm,
		MaxConcurrency: maxConcurrency,
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("healthy"))
}

func readinessCheck(w http.ResponseWriter, r *http.Request) {
	commandsMap.RLock()
	count := len(commandsMap.data)
	commandsMap.RUnlock()

	if count < cfg.MaxConcurrency {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ready"))
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

// Add to main() http handlers

// addCommandsHandler expects a JSON array, e.g.:
// [
//
//	{"command": "tail -f /dev/null", "write_file": "tail_output.txt"},
//	{"command": "sleep 99999", "write_file": "sleep_output.txt"}
//
// ]
func addCommandsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var requests []RequestBody
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requests); err != nil {
		http.Error(w, "Invalid JSON array", http.StatusBadRequest)
		return
	}

	for _, req := range requests {
		if req.Command == "" || req.WriteFile == "" {
			http.Error(w, "Each item must have 'command' and 'write_file'", http.StatusBadRequest)
			return
		}

		id := generateID()
		info := &commandInfo{
			ID:        id,
			Command:   req.Command,
			WriteFile: req.WriteFile,
			IsRunning: true,
		}

		commandsMap.Lock()
		commandsMap.data[id] = info
		commandsMap.Unlock()

		wg.Add(1)
		go runCommand(id, req.Command, req.WriteFile)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Commands accepted: %d\n", len(requests))))
}

// statusCommandsHandler returns a JSON array of all known commands:
// [
//
//	{ "id": "1679687111878627000-1", "command": "sleep 99999", "write_file": "sleep_output.txt", "is_running": true },
//	...
//
// ]
func statusCommandsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET is allowed", http.StatusMethodNotAllowed)
		return
	}

	commandsMap.RLock()
	defer commandsMap.RUnlock()

	responses := make([]commandInfo, 0, len(commandsMap.data))
	for _, info := range commandsMap.data {
		responses = append(responses, commandInfo{
			ID:        info.ID,
			Command:   info.Command,
			WriteFile: info.WriteFile,
			IsRunning: info.IsRunning,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responses); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

// deleteCommandHandler kills the process group for a single command, then removes it.
// Usage: DELETE /delete_command?id=1679687111878627000-1
func deleteCommandHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE is allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	commandsMap.Lock()
	info, ok := commandsMap.data[id]
	if !ok {
		commandsMap.Unlock()
		http.Error(w, "Command not found", http.StatusNotFound)
		return
	}

	if info.IsRunning && info.cmd != nil && info.cmd.Process != nil {
		pid := info.cmd.Process.Pid
		err := syscall.Kill(-pid, syscall.SIGKILL)
		if err != nil {
			fmt.Printf("Failed to kill process group for ID %s (PID=%d): %v\n", id, pid, err)
		}
		info.IsRunning = false
	}

	delete(commandsMap.data, id)
	commandsMap.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Command with ID %s deleted (killed if running).\n", id)))
}

// deleteAllCommandsHandler kills *all* commands (if running) and clears the map.
// Usage: DELETE /delete_all_commands
func deleteAllCommandsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Only DELETE is allowed", http.StatusMethodNotAllowed)
		return
	}

	commandsMap.Lock()
	count := len(commandsMap.data)
	for _, info := range commandsMap.data {
		if info.IsRunning && info.cmd != nil && info.cmd.Process != nil {
			pid := info.cmd.Process.Pid
			// Negative pid => kill entire process group
			_ = syscall.Kill(-pid, syscall.SIGKILL)
			// Mark not running in case the user calls status before it's deleted
			info.IsRunning = false
		}
	}
	// Clear all entries
	commandsMap.data = make(map[string]*commandInfo)
	commandsMap.Unlock()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("All commands (%d) deleted.\n", count)))
}

// runCommand starts the shell command in a new process group, so we can kill all descendants.
func runCommand(commandID, cmdStr, outFilename string) {
	defer wg.Done()

	fmt.Printf("Starting command [%s]: %s\n", commandID, cmdStr)

	cmd := exec.Command("bash", "-c", cmdStr)
	// On Unix-like systems, create a new process group
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	commandsMap.Lock()
	if info, ok := commandsMap.data[commandID]; ok {
		info.cmd = cmd
	}
	commandsMap.Unlock()

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating StdoutPipe for '%s': %v\n", cmdStr, err)
		setCommandFinished(commandID)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error starting command '%s': %v\n", cmdStr, err)
		setCommandFinished(commandID)
		return
	}

	file, err := os.OpenFile(outFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open file '%s': %v\n", outFilename, err)
		setCommandFinished(commandID)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		//timestamp := time.Now().Format(time.RFC3339)
		//_, wErr := writer.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, line))
		_, wErr := writer.WriteString(fmt.Sprintf("%s\n", line))

		if wErr != nil {
			fmt.Printf("Error writing to file '%s': %v\n", outFilename, wErr)
		}
		if fErr := writer.Flush(); fErr != nil {
			fmt.Printf("Error flushing writer for '%s': %v\n", outFilename, fErr)
		}
	}
	if scanErr := scanner.Err(); scanErr != nil {
		fmt.Printf("Error reading command '%s' output: %v\n", cmdStr, scanErr)
	}

	// Wait for command to finish (or remain stuck if never killed)
	if waitErr := cmd.Wait(); waitErr != nil {
		fmt.Printf("Command [%s] '%s' ended with error or signal: %v\n", commandID, cmdStr, waitErr)
	} else {
		fmt.Printf("Command [%s] '%s' finished successfully.\n", commandID, cmdStr)
	}

	setCommandFinished(commandID)
}

// setCommandFinished marks the command as not running in the map.
func setCommandFinished(commandID string) {
	commandsMap.Lock()
	if info, ok := commandsMap.data[commandID]; ok {
		info.IsRunning = false
	}
	commandsMap.Unlock()
}
