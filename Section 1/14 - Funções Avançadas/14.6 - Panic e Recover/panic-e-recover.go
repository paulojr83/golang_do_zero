package main

import "fmt"

func recuperarExeccao()  {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}
func alunoAprovado(n1, n2 float32) bool  {
	defer recuperarExeccao()

	media := ( n1 + n2 ) / 2

	if  media > 6 {
		return true
	}else if media < 6 {
		return false
	}
  	panic("A MÉDIA É EXATAMENTE 6!")
}
func main() {
	fmt.Println(alunoAprovado(10, 7))
	fmt.Println("Pós execução")
}
