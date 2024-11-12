package main

import "fmt"

func main() {

	/**
		Defer é um statment que for a execução de um trecho por último.
	Ajuda a fechar conexões.
	**/
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
}