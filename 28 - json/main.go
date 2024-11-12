package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
json:.... = essa anotação são chamadas de tags, são usadas para serializar e deserializar objetos.
**/
type Pessoa struct {
	Nome string `json:"nome"`
	Idade int `json:"idade"`
	Oculto int `json:"-"`
}


func main() {
	// JSON
	// Para trabalhar com JSON, você pode usar a biblioteca padrão encoding/json.
	// Para converter um objeto para JSON, use a função json.Marshal

	pessoa := Pessoa{
		Nome: "João",
		Idade: 30,
		Oculto: 1,
	}

	content, err := json.Marshal(pessoa)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))


	// Para converter JSON para um objeto, use a função json.Unmarshal
	var decodedPerson Pessoa
	jsonString := []byte(`{"nome":"João","idade":30, "Oculto": "1"}`)

	err = json.Unmarshal(jsonString, &decodedPerson)

	if err != nil {
		panic(err)
	}

	fmt.Println(decodedPerson.Nome)

	err = json.NewEncoder(os.Stdout).Encode(pessoa)
	if err != nil {
		panic(err)
	}
}