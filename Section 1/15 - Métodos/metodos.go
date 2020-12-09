package main

import "fmt"

type usuario struct {
	nome 	string
	idade	uint8
}

func (u usuario) salvar()  {
	fmt.Printf("Salvando usuário:%s\n", u.nome)
}

func (u usuario) maiorDeIdade() bool  {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario()  {
	u.idade++
}

func main() {
	u := usuario{
		nome:  "Paulo",
		idade: 37,
	}

	u.salvar()
	u.fazerAniversario()
	ehMaiorDeIdade := u.maiorDeIdade()
	fmt.Println(u, ehMaiorDeIdade)
}
