package main

import "errors"

func Exercicio2(valor ...float64) (float64, error) {
	var total float64
	var divisao = len(valor)
	for i := 0; i < len(valor); i++ {
		if valor[i] < 0 {
			return 0, errors.New("nao pode ser numero negativo")
		} else {
			total += valor[i]
		}
	}
	return total / float64(divisao), nil
}
