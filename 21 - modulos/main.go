package main

import (
	"fmt"

	"github.com/diegosilva19/go-curso/matematica"
	"github.com/google/uuid"
)

/**
go mod init github.com/diegosilva19/go-curso
go get golang.org/x/exp/constraints
go mod tidy    -- baixa e equaliza pacotes baixado e não utilizados

* Para exportar funções/ variaveis / structs / propriedades de structs/ funções de struct  de um pacote, a primeira letra da função deve ser maiúscula
**/
func main() {

	fmt.Println("Soma de 1 + 1 é", matematica.Soma(1, 1))
	fmt.Println(uuid.New())
}