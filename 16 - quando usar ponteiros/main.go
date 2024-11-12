package main

//funcao sem uso de referencia - fiel ao valor que foi passado
func soma(a * int, b int) int {
	*a = 60
	return *a+b
}

func main() {


	numeroFixo1 := 10
	numeroFixo2 := 20
	soma(&numeroFixo1, numeroFixo2)

	println(numeroFixo1)

}
