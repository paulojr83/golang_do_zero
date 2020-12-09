package main

import "fmt"

func funcao1()  {
	fmt.Println("Função 1")
}

func funcao2()  {
	fmt.Println("Função 2")
}

func alunoAprovado(n1, n2 float32) bool  {
	defer fmt.Println("Média calculada. Resultado será retornada.")
	fmt.Println("Entradando na função para verificar se o aluno está aprovado.")

	media := ( n1 + n2 ) / 2

	if  media >= 6 {
		return true
	}

	return false
}
func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunoAprovado(10, 2))
}
