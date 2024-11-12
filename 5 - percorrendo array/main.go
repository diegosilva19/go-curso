package main

import "fmt"

type ID int

var (
	b bool
	c int
	d string = "nome qualquer coisa inicializando"
	e float64
	f ID = 1
)


func main() {

	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray[len(meuArray) - 1])

	for idx, valor := range meuArray {
		fmt.Printf("O valor do índíce %d, é %d\n", idx, valor)
	}
}
