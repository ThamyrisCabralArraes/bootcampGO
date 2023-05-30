package main

import (
	"fmt"
	"os"
)

func exercicio1() {
	leu, err := os.Open("customers.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(leu)
}
