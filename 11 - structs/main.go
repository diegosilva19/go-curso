package main

import "fmt"

type Client struct {

	Nome string
	Idade int
	Ativo bool
}

func main() {

	client1 := Client{
		Nome: "Diego",
		Idade: 34,
		Ativo: true,
	}

	client2 := Client{
		Nome: "João",
		Idade: 10,
		Ativo: false,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %v\n", client1.Nome, client1.Idade, client1.Ativo )
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %v\n", client2.Nome, client2.Idade, client2.Ativo )
}
