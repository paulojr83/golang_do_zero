package main

import "fmt"

func somar(n1 int8, n2 int8) int8  {

	return n1 + n2
}

func caluclosMatematicos(n1, n2 int8) (int8, int8)  {
	 soma := n1 + n2
	 subtracao := n1 - n2
	 return soma, subtracao
}
func main() {
	soma := somar(20, 30)
	fmt.Println(soma)
	fmt.Println(somar(2,9))

	var f = func(txt string) string {
		return txt
	}

	resultado := f("Outra função")
	fmt.Println(resultado)
	s, sub := caluclosMatematicos(10, 29)
	fmt.Printf("Soma:%d | Subritação: %d\n",s, sub)
}
