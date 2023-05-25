package main

import "fmt"

func Exercicio2() {
	clienteMaria := map[string]int{"idade": 30, "empregago": 1, "ano de atividade": 2, "salario": 50.000}

	if clienteMaria["idade"] > 20 && clienteMaria["empregago"] == 1 && clienteMaria["ano de atividade"] > 1 {
		fmt.Println("Pode fazer um emprestimo")
		if clienteMaria["salario"] > 100.000 {
			fmt.Println("Sem juros")
		} else {
			fmt.Println("Com Juros")
		}
	} else {
		fmt.Println("Não pode fazer um emprestimo")
	}
}
