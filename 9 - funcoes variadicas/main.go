package main

import (
	"errors"
	"fmt"
)

func main() {


	//------------------------------------------------------------------
	// soma que devolve um erro se o resultado por par
	valor, errorMessage := sumError(1, 4, 2)

	if errorMessage != nil {
		fmt.Println(errorMessage)
	}
	fmt.Println(valor)
}

// funcoes variadicas são funçẽso com N parametros - é preciso iterar e executar a ação
func sumError(numeros ...int) (int, error) {

	result := 0
	for _, numero := range numeros {
		result += numero
	}

	if result % 2 == 0 {
		return 0, errors.New("só entrego a soma se o resultado dos itens não for par exemplo -> 1 + 4 +2")
	}
	return result, nil
}
