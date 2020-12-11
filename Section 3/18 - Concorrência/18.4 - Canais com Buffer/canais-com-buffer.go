package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Golang 1"
	canal <- "Olá Golang 2"

	mensagem := <- canal
	fmt.Println(mensagem)
}
