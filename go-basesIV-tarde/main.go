package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// exercicio1()
	// exercicio2()

	//json exemplo 1
	type product struct {
		Name  string
		Preco int
	}

	p := product{
		Name:  "thamy",
		Preco: 1,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	// json exemplo2
	jsonData2 := `{"Name": "MacBook Air", "Preco": 900}`

	var resposta product
	if erro := json.Unmarshal([]byte(jsonData2), &resposta); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(resposta)

	// mesma coisa que o MArshall faz
	encode := product{
		Name:  "carlos",
		Preco: 2,
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(encode)

	// mesma coisa do Unmarshal
	type x struct {
		A string
		B int
		C bool
	}

	jsonData3 := `{
		"A": "Olá Mundo",
		"B": 5,
		"C": true}`

	z := x{}

	buff := bytes.NewBuffer([]byte(jsonData3))
	decoder := json.NewDecoder(buff)
	decoder.Decode(&z)

	fmt.Println(z)

}
