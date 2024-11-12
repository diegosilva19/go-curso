package main

import "fmt"

func main() {

	salarios := map[string]int{"Diego": 300, "João": 200};

	//fmt.Println(salarios)

	delete(salarios, "Diego")
	//fmt.Println(salarios, salarios["pessoa2"])



	salarios2 := make(map[string]int)
	salarios2["teste"] = 200;

	//fmt.Println(salarios2)

	for nome, salarioInfo := range salarios {
		fmt.Printf("O salario de %s é %d \n", nome, salarioInfo)
	}

	// ignorar o nome
	for _, salarioInfo := range salarios {
		fmt.Printf("O salario é %d \n", salarioInfo)
	}
}