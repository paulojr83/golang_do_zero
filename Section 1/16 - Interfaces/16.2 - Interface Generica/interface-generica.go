package main

import (
	"fmt"
	"reflect"
)

func generica(interf interface{})  {
	fmt.Printf("Tipo:%s || Valor:%v\n", reflect.TypeOf(interf), interf)
}
func main() {
	generica("Teste")
	generica(10)

}
