package main

import "fmt"

func main() {

	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	usuario2 := map[int]map[string]string {
		1: {
			"nome": "Paulo",
		},
	}
	fmt.Println(usuario2[1])

}
