package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type KubePod struct {
	Name          string `json:"-"`
	Namespace     string `json:"-"`
	ClusterName   string `json:"-"`
	AppName       string `json:"-"`
	DeployVersion string `json:"-"`
	PodNameHash   string `json:"-"`
	LogCommandStr string `json:"command"`
	WriteFile     string `json:"write_file"`
}

type RequestBodyList struct {
	ClusterName   string `json:"cluster_name"`
	NameSpace     string "json:namespace"
	LogDir        string `json:"log_dir"`
	SeparatedFile bool   `json:"separated_file"`
}

func CreateList(w http.ResponseWriter, r *http.Request) {

	var payload []RequestBodyList
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON array", http.StatusBadRequest)
		return
	}

	var responses []KubePod

	for _, itemList := range payload {
		clusterName := itemList.ClusterName
		namespace := itemList.NameSpace

		cmdStr := "kubectx " + itemList.ClusterName + " && kubectl get pods -n " + namespace + " -o json | jq -r '.items[] | select(.metadata.name | contains(\"" + namespace + "\")) | .metadata.name'"
		cmd := exec.Command("bash", "-c", cmdStr)
		stdout, err := cmd.StdoutPipe()

		if err != nil {
			http.Error(w, "Failed to execute command", http.StatusInternalServerError)
			return
		}

		if err := cmd.Start(); err != nil {
			http.Error(w, "Failed to start command", http.StatusInternalServerError)
			return
		}
		scanner := bufio.NewScanner(stdout)

		for scanner.Scan() {
			line := scanner.Text()

			partsPodName := strings.Split(line, "-")

			if len(partsPodName) == 3 {

				appName := partsPodName[0]
				deployVersion := partsPodName[1]
				podNameHash := partsPodName[2]

				podInfo := KubePod{
					Namespace:     namespace,
					AppName:       appName,
					ClusterName:   clusterName,
					DeployVersion: deployVersion,
					PodNameHash:   podNameHash,
					LogCommandStr: "kubectl logs -f " + line,
					WriteFile:     itemList.LogDir + "/" + appName + "-" + deployVersion + "-" + podNameHash + ".log",
				}

				if !itemList.SeparatedFile {
					podInfo.WriteFile = "logs/" + appName + ".log"
				}
				responses = append(responses, podInfo)
			}

		}

		if waitErr := cmd.Wait(); waitErr != nil {
			fmt.Printf("Command '%s' ended with error or signal: %v\n", cmdStr, waitErr)
		} else {
			fmt.Printf("Command '%s' finished successfully.\n", cmdStr)
		}

	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responses); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}

}
