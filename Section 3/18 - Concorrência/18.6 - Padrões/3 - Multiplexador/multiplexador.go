package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Ola, Golang!"), escrever("Vamos Golang!"))
	for i:=0; i< 10;i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalDeSaida :=make(chan string)

	go func() {
		for  {
			select {
			case mensagem := <- canalEntrada1:
				canalDeSaida <-mensagem
			case mensagem := <-canalEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}
func escrever(texto string) <-chan string{
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprint(texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}