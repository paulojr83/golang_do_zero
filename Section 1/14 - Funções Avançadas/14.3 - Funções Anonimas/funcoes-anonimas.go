package main

import "fmt"

func main() {
	func(versao float32){
		fmt.Printf("Olá Golang!! %.2f\n", versao)
	}(1.15)

	retorno := func(versao float32) string {
		return fmt.Sprintf("Olá Golang!! %.2f\n", versao)
	}(1.15)

	fmt.Println(retorno)
}
