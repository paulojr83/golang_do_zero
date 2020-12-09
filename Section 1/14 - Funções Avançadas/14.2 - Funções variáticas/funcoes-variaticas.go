package main

import "fmt"

func somar(numeros ...int)int  {
	total := 0
	for _, numero := range numeros {
		total +=numero
	}

	return total
}

func escrever(texto string, numeros ...float32)  {
	fmt.Println()
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}
func main() {

	fmt.Println("Funções variáticas")
	fmt.Println(somar(1,2, 3, 4, 4, 5, 6,7, 76, 8,87 ,87))
	escrever("Olá Go", 1.15)
}
