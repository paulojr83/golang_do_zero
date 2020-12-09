package main

import (
	"fmt"
	"time"
)

func main() {
 	i :=0

	for i < 5 {
		time.Sleep(time.Second)
		i++
		fmt.Printf("Incrementando i %d\n", i)
	}

	for j := 0; j < 1; j++ {
		time.Sleep(time.Second)
		fmt.Printf("Incrementando j %d\n", j)
	}

	nomes := []string{"João", "Pedro", "Zé"}
	for _, n := range nomes {
		fmt.Println(n)
	}

	for index, letra := range "PALAVRA" {
		fmt.Println(index, string(letra))
	}

	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Printf("Chave:%s | Valor:%s\n", chave, valor)
	}
	for {

		fmt.Println("Loops infinito!!!")
		time.Sleep(time.Second)
	}

}
