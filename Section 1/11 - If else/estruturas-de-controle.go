package main

import "fmt"

func main() {
	fmt.Println("Estrutura de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	}else {
		fmt.Println("Menor ou igual a 15")
	}

	// Variavel é valida somente no scopo do if/else
	if outroNumero := numero; outroNumero > 0 {
		fmt.Printf("Númerp é maior que zero: %d\n", outroNumero)
	}else {
		fmt.Printf("Número é menor que zero: %d\n", outroNumero)
	}

}
