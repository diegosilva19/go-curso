package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, error := os.Create("file.txt")

	if error != nil {
		panic(error)
	}

	size, error := file.Write([]byte("Hello, World!"));
	//size, error := file.WriteString("Hello, World!");

	if error != nil {
		panic(error)
	}
	fmt.Printf("arquivo criado com sucesso! tamanho : %d bytes", size)

	file.Close()


	// lendo
	arquivo, error := os.ReadFile("file.txt")

	if error != nil {
		panic(error)
	}
	fmt.Println(string(arquivo))

	// lendo o arquivo linha a linha
	arquivoLinha, error := os.Open("file.txt")
	if error != nil {
		panic(error)
	}
	reader := bufio.NewReader(arquivoLinha)

	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println("n: ", n, " - ")
		fmt.Println(string(buffer[:n]))
	}

	//removendo
	error = os.Remove("file.txt")
	if error != nil {
		panic(error)
	}
}