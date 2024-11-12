package main

import "fmt"

type Conta struct {
	NomePessoa string
	saldo float64
}

// sem o * no parametro não altera a referencia da variável
func (c* Conta) simular(valor float64) float64 {
	c.saldo += valor;
	return c.saldo
}


func NewConta(nomePessoa string, saldoConta float64) *Conta {
	return &Conta{
		NomePessoa: nomePessoa,
		saldo: saldoConta,
	}
}

func main() {

	// map de chave valores pre definidos
	contas :=map[int]Conta{}
	contas[0] = Conta{
		NomePessoa: "Diego Silva",
		saldo: 200.23,
	}
	contas[1] = Conta{
		NomePessoa: "Diego Silva",
		saldo: 200.23,
	}

	
	// append em um objetivo infinito
	contasInfinitas := []Conta{}
	contasInfinitas = append(contasInfinitas, Conta{
		NomePessoa: "Pessoa 1",
		saldo: 10.00,
	})

	contasInfinitas = append(contasInfinitas, *NewConta("Pessoa 2", 20.00))
	contasInfinitas = append(contasInfinitas, *NewConta("Pessoa 3", 20.00))
	

	for index, conta := range contasInfinitas {
		fmt.Printf("Nome %s, saldo %v, index %v\n", conta.NomePessoa, conta.saldo, index)
	}
	fmt.Printf("Total de contas %v\n", len(contasInfinitas))
}