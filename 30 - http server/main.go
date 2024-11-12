package main

import "net/http"

func main() {

	http.ListenAndServe(":8080", nil)
}

func BuscaCep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca Cep"))
}