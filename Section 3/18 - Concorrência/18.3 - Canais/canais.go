package main

import (
	"fmt"
	"time"
)

func main() {

	canal := make(chan string)
	go escrever("Ola Golang", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <- canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do Programa")

}


func escrever(texto string, canal chan string)  {
	for i :=0; i < 5; i++  {
		canal <- fmt.Sprintf("%d) %s", i + 1, texto)
		time.Sleep(time.Second * 2)
	}
	close(canal)
}
