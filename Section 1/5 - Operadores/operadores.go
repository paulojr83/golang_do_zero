package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Subtração => ", a-b)
	fmt.Println("Divisão => ", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo => ", a%b)

	// bitwise,
	fmt.Println("E =>", a&b)
	fmt.Println("Ou =>", a|b)
	fmt.Println("Xor =>", a^b)

	c := 3.0
	d := 2.0

	// Outras operações usando math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))

	var numero1 int16 = 10
	var numero2 int8 =  10
	fmt.Println(numero1, numero2)
	/*
	var soma = numero1 + numero2
	fmt.Println(soma) // Não é possivel fazer esse tipo de operação, pois são tipos diferentes
	 */

	// OPERADORES LÓGICOS
	fmt.Println("String:", "Banana" == "Banana")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)
	fmt.Println("<=", 3 <= 2)
	fmt.Println(">=", 3 >= 2)

	d1 := time.Unix(0,0)
	d2 := time.Unix(0,0)
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{ "João" }
	p2 := Pessoa{ "João" }

	fmt.Println("Pessoas:", p1 == p2)

	// OPERADORES UNÁRIOS
	x := 1
	y := 2

	// apenas postfix
	x++ // x +=1 ou x = x + 1
	y-- // x -=1 ou y = y - 1
	fmt.Println(x)
	fmt.Println(y)

	x *=3
	fmt.Println(x)
	x /=10
	fmt.Println(x)
	x %=3
	fmt.Println(x)

	// não existe ternário no Go
	// x == 3 ? true : false

}
