package main

import (
	"io"
	"net/http"
)

func main() {
	// Chamada HTTP
	// Para fazer chamadas HTTP, você pode usar a biblioteca padrão net/http.
	req, err := http.Get("http://google.com")

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))

}
