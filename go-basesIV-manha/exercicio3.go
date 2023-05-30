package main

import (
	"fmt"
)

func exercicio3(s float64) {
	if s < 15.000 {
		err := fmt.Errorf("erro3: O salário digitado não está dentro do valor mínimo")
		fmt.Println(err)
		return
	}
	fmt.Println("mensagem 3: Necessário pagamento de imposto.")
}
