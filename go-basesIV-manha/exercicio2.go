package main

import (
	"errors"
	"fmt"
)

func exercicio2(s float64) {
	if s < 15.000 {
		fmt.Println(errors.New("erro: O salário digitado não está dentro do valor mínimo"))
		return
	}
	fmt.Println("mensagem 2: Necessário pagamento de imposto.")
}
