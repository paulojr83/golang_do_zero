package main

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

func escreverArea(f Forma)  {
	fmt.Printf("A área da forma é %.2f\n", f.area())
}

type retangulo struct {
	altura		float64
	largura		float64
}

type circulo struct {
	raio		float64
}

func (r retangulo) area() float64  {
	return r.altura * r.largura
}

func (c circulo) area() float64  {
	return math.Pi * math.Pow( c.raio, 2 )
}

func main() {
	r := retangulo{
		altura:  10,
		largura: 10,
	}

	escreverArea(r)

	c := circulo{
		raio:  10,
	}
	escreverArea(c)

}
