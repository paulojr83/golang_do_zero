package main

import "fmt"

type pessoa struct {
	nome		string
	sobrenome	string
	idade		uint8
	altura		uint8
}

type esdudante struct {
	pessoa
	curso 		string
	faculdade	string
}

func main() {
	p1 := pessoa{
		nome:      "Paulo",
		sobrenome: "Porto",
		idade:     37,
		altura:    183,
	}
	fmt.Println(p1)

	e1:= esdudante{
		pessoa:p1,
		curso:     "Go lang",
		faculdade: "Udemy",
	}
	fmt.Println(e1.nome)
}
