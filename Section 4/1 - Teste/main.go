package main

import (
	"enderecos/enderecos"
	"fmt"
)

func main() {
	endereco := enderecos.TipoDeEndereco("Rua das bananeiras")
	fmt.Println(endereco)

}
