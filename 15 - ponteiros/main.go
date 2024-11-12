package main

import "fmt"


type Endereco struct {
	Logradouro string
	Number int
	Cidade string
	Estado string

}

type Pessoa interface {
	Desativar()
}

type Client struct {

	Nome string
	Idade int
	Ativo bool
	Address Endereco
}

func (c Client) Desativar() {
	c.Ativo = false
	fmt.Printf("Desativando cliente -> %s\n", c.Nome)
}


func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	client1 := Client{
		Nome: "Diego",
		Idade: 34,
		Ativo: true,
	}

	var clientPointer *Client = &client1
	clientPointer.Idade = 20
	println("endereço de memória -> ", &client1, clientPointer)// pegar endereço de memoria
	fmt.Printf(" Idade %d\n", clientPointer.Idade)


	a:= 10
	var ponteiro *int = &a
	*ponteiro = 25
	b :=&a
	*b = 50 // atribuir valor ao ponteiro sempre colocar * antes

	println(b) // printa endereço de memória
	println(*b) // direference - printa conteúdo endereço de memória

}
