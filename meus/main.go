package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Agencia struct {
	nome   string
	numero int16
}

type Operacao struct {
	id    string
	nome  string
	valor float64
	data  string
}

type Conta struct {
	saldo     float64
	agencia   Agencia
	operacoes []Operacao
}

var (
	ct Conta
	ag Agencia
)

func main() {

	ag = Agencia{
		"agencia exemplo",
		130,
	}
	ct = Conta{
		15.30,
		ag,
		nil,
	}

	reader := bufio.NewReader(os.Stdin)

	operations := map[string]int{"Débito": 1, "Crédito": 2}
	answer := Question("Informe a operação que deseja fazer.\n\n", operations, reader)

	fmt.Print("Resposta -> " + answer + "\n")
}

func Question(question string, options map[string]int, reader *bufio.Reader) string {

	fmt.Print(question)
	var expectedValue string
	for {
		var possibleOptions []string
		for optionName, optionValue := range options {
			fmt.Print("(" + strconv.Itoa(optionValue) + ") : " + optionName + "\n")
			possibleOptions = append(possibleOptions, strconv.Itoa(optionValue))
		}

		content, _ := reader.ReadString('\n')
		inputContent := strings.TrimRight(content, "\n")

		for _, opt := range possibleOptions {
			if strings.Compare(opt, inputContent) == 0 {
				expectedValue = inputContent
			}
		}

		if expectedValue == "" {
			fmt.Print("opção inválida tente denovo\n\n")
		} else {
			break
		}
	}

	return expectedValue
}
