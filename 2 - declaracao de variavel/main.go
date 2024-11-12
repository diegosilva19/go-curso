package main

const a = "hello, world!"

var b bool
var (
	c int
	d string = "nome qualquer coisa inicializando"
	e float64
)

func main() {

	f := 10 // inferencia de tipo automatica invés de declarar assim já cria direto
	f = 200
	println(a, b, c, d, e, f)
}
