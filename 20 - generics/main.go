package main

import "fmt"

// Generics a partir da versão 1.18

func SomaInteiro(m map[string]int) int {
	var soma int;
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat64(m map[string]float64) float64 {
	var soma float64;
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaGenerico[T int | float64](m map[string]T) T {
	var soma T;
	for _, v := range m {
		soma += v
	}
	return soma
}
/*********************************** Constraint invés de declarar tipo a tipo no generic ***********************************/
type MyNumber int


// quando estou forçando o meu tipo neste ponto embora MyNumber seja int como o generic que tem int | float64 ainda assim eu preciso adicionar o ~ para que ele aceite
type Number interface {
	~int | float64
}

func SomaGenericoConstraint[T Number](m map[string]T) T {
	var soma T;
	for _, v := range m {
		soma += v
	}
	return soma
}

/*********************************** 
 constraint "comparable" é um tipo de interface que permite que o tipo seja comparável - ou seja ao utilizar a função a entrada será checada
	só funciona com a igudalidade ==
***********************************/
func Compare[T comparable](a T, b T) bool {
	return a == b
}

func main() {

	/*intMap := map[string]int{"Diego": 1000, "Cleiton": 500, "Robson": 200}
	floatMap := map[string]float64{"Diego": 1000.10, "Cleiton": 500.20, "Robson": 200.30}
	println(SomaInteiro(intMap))
	println(SomaFloat64(floatMap))
	println(SomaGenerico(intMap))
	println(SomaGenerico(floatMap))
	
	// quando estou forçando o meu tipo neste ponto embora MyNumber seja int como o generic que tem int | float64 ainda assim eu preciso adicionar o ~ para que ele aceite
	myNumberMap := map[string]MyNumber{"Diego": 1000, "Cleiton": 500, "Robson": 200}

	fmt.Println(SomaGenericoConstraint(myNumberMap))*/
	fmt.Println(Compare(10.10, 1))

}