package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, subtracao int)  {
	soma = n1 + n2
	subtracao = n1 - n2

	return
}

func main() {

	fmt.Println("Funções Avançadas")

	soma, subtracao :=calculosMatematicos(10, 2)
	fmt.Printf("Soma:%d | Subtração:%d\n", soma, subtracao)
}
