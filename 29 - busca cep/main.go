package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)


type Endereco struct {
	Cep string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade string `json:"unidade"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf string `json:"uf"`
	Estado string `json:"estado"`
	Regiao string `json:"regiao"`
	Ibge string `json:"ibge"`
	Gia string `json:"gia"`
	Ddd string `json:"ddd"`
	Siafi string `json:"siafi"`
}
func main() {

	for _, cep := range os.Args[1:] {

		url := "https://viacep.com.br/ws/" + cep + "/json"
		req, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer request reposta: %v\n", err)
		}
		defer req.Body.Close()
		response, err := io.ReadAll(req.Body);

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler reposta: %v\n", err)
		}

		var endereco Endereco;
		err = json.Unmarshal(response, &endereco)

		if err != nil {
			println(cep)
		}

		jsonString, err := json.Marshal(&endereco);

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer marshal: %v\n", err)
			panic(1)
		}

		//go writeOnDist("searchs.json", string(jsonString))
		//go writeOnDist("address.txt", "CEP: " + endereco.Cep + " - " + endereco.Logradouro + " - " + endereco.Bairro + " - " + endereco.Localidade + " - " + endereco.Uf)
		
		writeOnDist("searchs.json", string(jsonString))
		writeOnDist("address.txt", "CEP: " + endereco.Cep + " - " + endereco.Logradouro + " - " + endereco.Bairro + " - " + endereco.Localidade + " - " + endereco.Uf)

		fmt.Println("\nResultado da request -> " + string(jsonString))
		
	}
}

func writeOnDist(fileName string, message string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}

		defer file.Close()
		totalByes, err := file.WriteString(string(message)+"\n")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever arquivo: %v\n", err)
		}
		fmt.Printf("Write %d bytes", totalByes)
}