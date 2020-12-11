package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	go escrever("Ola Golang 1") // gorouting
	escrever("Ola Golang 2")
}

func escrever(texto string)  {
	for  {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
