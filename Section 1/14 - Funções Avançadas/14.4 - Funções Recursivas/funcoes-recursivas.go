package main

import "fmt"

func fibonate(posicao uint) uint {
	if(posicao <= 1 ) {
		return posicao
	}
	return fibonate(posicao - 2) + fibonate(posicao -1)
}
func main() {
	fmt.Println("Funções Recursivas")
	posicao := uint(10)

	for i:= uint(0); i <= posicao; i++ {
		fmt.Println(fibonate(i))
	}

}
