package main

import "fmt"

func main() {

	s := []int{2, 4, 6, 8, 10}

	var itensFrenteTras []int
	itensFrenteTras = s[2:]

	itensTrasFrente := s[:0]

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), itensTrasFrente) // ...2 posição para trás
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), itensFrenteTras) // 2... posição para frente

	//cuidado !!!! toda vez que se da uma append ele copia o slice / array para outro com o dobro do tamanho original
	s = append(s, 100)

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:]) // 2... posição para frente

}
