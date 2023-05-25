package main

import "fmt"

func Exercicio1() {
	var nome string = "Thamyris"

	fmt.Println("quantidade de letras ", len(nome))

	for _, letra := range nome {
		fmt.Printf("letras %c\n ", letra)
	}
}
