package main

import (
	"fmt"
	"time"
)

func main() {

	canal1, canal2 := make(chan string), make(chan string)

	var index1, index2 int
	go func() {
		for{
			index1++
			time.Sleep(time.Millisecond * 500)
			canal1 <- fmt.Sprintf("%d) Canal 1", index1)
		}

	}()

	go func() {
		for {
			index2++
			time.Sleep(time.Second * 2)
			canal2 <- fmt.Sprintf("%d) Canal 2", index2)
		}

	}()

	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)

		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}

}
