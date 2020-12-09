package main

import "fmt"

type usuario struct {
	name 		string  `json:"name"`
	idade		uint8   `json:"idade"`
	endereco	endereco
}

type endereco struct {
	logradouro 	string
	numero		uint16
}

func main() {

	var u usuario
	u.name = "Paulo"
	u.idade = 37
	fmt.Println(u)
	endereco := endereco{}

	u2 := usuario{ "Priscilla", 31,endereco, }
	fmt.Println(u2)

	endereco.logradouro ="Rua A"
	u3 := usuario{name: "Porto", endereco:endereco}
	fmt.Println(u3)
	fmt.Println("Arquivo struts")
}
