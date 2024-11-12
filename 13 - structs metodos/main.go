package main

import "fmt"


type Endereco struct {
	Logradouro string
	Number int
	Cidade string
	Estado string

}
type Client struct {

	Nome string
	Idade int
	Ativo bool
	Address Endereco
}

// aqui por não ter o * gera erro de referencia na hora de definir o comportamento mas no próximo passo será corrigido
func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("Desativando cliente -> %s\n", c.Nome)
}

func main() {

	client1 := Client{
		Nome: "Diego",
		Idade: 34,
		Ativo: true,
	}

	client1.Address.Cidade = "Diadema"
	client1.Desativar()

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %v\n", client1.Nome, client1.Idade, client1.Ativo )
}
