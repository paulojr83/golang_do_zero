package main

import (
	"log"
	"net/http"
)

func home(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("Olá Golang!!"))
}
func main() {
	// HTTP é um protocolo de comunicação - Base da comunicação WEB

	// Clinte (Faz Requisição ) - Servidor ( Processa requisição e envia resposta )
	// Request - Response

	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE, PATCH

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
