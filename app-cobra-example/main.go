package app_cobra_example

import "fmt"

var mensagem string

func main() string {
	mensagem = "okok"
	fmt.Printf("ola mundo denovo")
	return mensagem
}


func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
