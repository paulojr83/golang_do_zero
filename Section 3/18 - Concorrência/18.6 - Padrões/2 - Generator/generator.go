package main

import (
	"fmt"
	"time"
)

func main() {
 	canal := escrever("Olá Golang!")
	for i := 0; i < 10; i++{
		fmt.Printf("%d) %s\n", i+1, <-canal)
	}
}

func escrever(texto string) <-chan string{
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprint(texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}