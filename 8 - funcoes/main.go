package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
	var (
		result int
		isPar bool
	)*/

	//------------------------------------------------------------------
		// soma simples
	//fmt.Println(sum(1,2))

	//------------------------------------------------------------------
		// soman com mais de um retorno dizendo se é par ou não
	//result, isPar = sumPar(2,2)
	//fmt.Printf("Somatoria é %d, numero par: %v\n", result, isPar)


	//------------------------------------------------------------------
	// soma que devolve um erro se o resultado por par
	valor, errorMessage := sumError(1,2)

	if errorMessage != nil {
		fmt.Println(errorMessage)
	}
	fmt.Println(valor)
}

// função que devolve erro, Go não possuí exception
func sumError(a, b int) (int, error) {

	result := a+b

	if result % 2 == 0 {
		return 0, errors.New("só entrego a soma se o resultado da mesma não for par exemplo -> 1+2")
	}
	return result, nil
}

func sumPar(a, b int) (int, bool) {

	var isPar bool
	result := a+b

	if result % 2 == 0 {
		isPar = true
	}
	return result, isPar
}

func sum(a int, b int) int {
	return a+b
}

//parametros iguais declara o tipo só um avez
func sumIguais(a, b int) int {
	return a+b
}


