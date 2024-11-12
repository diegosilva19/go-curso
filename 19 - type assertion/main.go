package main

import (
	"fmt"
	"strconv"
)

func main() {

	var x interface{} = 10;
	var y interface{} = "Hello, World!";
	showType(x)
	
	//print(y)
	//print(y.(string)) // forçar tipo

	valor, isValidAssertion := y.(int)
	
	// quando não tratado gera o panic error se a asserção falhar
	//valor := y.(int)

	if isValidAssertion {
		println("deu bom é inteiro" , strconv.Itoa(valor))
	} else {
		println("Não é um inteiro")
	}
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variavel é: %T e o valor é %v\n", t, t)
}