package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome	string `json:"nome"`
	Raca	string	`json:"raca"`
	Idade	uint	`json:"idade"`
}

func main() {
	 c := cachorro{
		 Nome:  "Tobby",
		 Raca:  "Vira latas",
		 Idade: 10,
	 }

	 cJson, err := json.Marshal(c)
	 if err != nil {
	 	log.Fatal(err)
	 }

	 //fmt.Println(string(json))
	fmt.Println(bytes.NewBuffer(cJson))

	c2 := map[string] string {
		"nome":  "Tobby 2",
		"raca":  "Vira latas",
	}

	jsonC2, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(jsonC2))
}
