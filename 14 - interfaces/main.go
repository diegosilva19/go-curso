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

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

// métodos são atreladas a structs a artir da definição inicial do objeto (c Client) diferente de uma classe que visualmente
// na estrutura tudo é agregado, no Go é o contrário assim com oa interface, o fato da entidade ter o Desativar, implicitamente ela já é do tipo "Pessoa"
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

	client1.Address.Cidade = "Diadema"

	empresa := Empresa{Nome: "empresa teste"}

	Desativacao(client1)
	Desativacao(empresa)

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %v\n", client1.Nome, client1.Idade, client1.Ativo )
}
