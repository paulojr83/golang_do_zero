package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJson := `{"nome":"Tobby","raca":"Vira latas","idade":10}`

	var c cachorro
	if erro:= json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c.Nome)

	cachorro2EmJson := `{"nome":"Tobby 2","raca":"Vira latas"}`

	c2:= make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
