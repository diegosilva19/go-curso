package main

import (
	"errors"
	"fmt"
)

func main() {


	//------------------------------------------------------------------
	// soma que devolve um erro se o resultado por par
	// na função anonima quando o retorno é de multiplos valores o que é atribuido a variável é uma instancia para rodar - execução - como no js
	callback := func() (int, error) {
		resultado, errorOccur := sumError(1, 4, 2)

		if errorOccur == nil {
			return resultado * 2, nil
		}
		return resultado, errorOccur
	}

	valorPuro := func() int {
		resultado, _ := sumError(1, 4, 2)
		return resultado * 2
	}

	result, errorMessage := callback()
	if errorMessage != nil {
		fmt.Println(errorMessage)
	}
	fmt.Println(valorPuro(), result)
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
