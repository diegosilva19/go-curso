package main

import "fmt"

type Conta struct {
	Titular Pessoa
	Saldo float64
}

type Pessoa struct {
	Nome string
	Idade int
}


func main() {

	contas := []Conta{};
	contasMapa := map[int]Conta{};

	contas = append(contas, Conta{Pessoa{Nome: "Diego", Idade: 35}, 20.50})
	contasMapa[1] = Conta{Pessoa{Nome: "Diego", Idade: 35}, 20.50}

	fmt.Println("Looping for com iterador")
	for _, contas := range contas {
		fmt.Printf("Titular: %s, Saldo: %0.1f", contas.Titular.Nome, contas.Saldo)
	}

	fmt.Println("Looping for com contador")
	// This is a loop
	for i := 0; i < len(contasMapa); i++ {

		fmt.Printf("Titular: %s, Saldo: %0.1f", contas[i].Titular.Nome, contas[i].Saldo)
		println(i)
	}

	//semelhante ao while
	i :=0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//for {
	//	fmt.Println("Looping infinito")
	//}
}