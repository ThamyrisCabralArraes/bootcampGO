package main

func Exercicio1(valor float64) float64 {
	var valorFinal float64

	if valor == 50.000 {
		valorFinal = valor * 1.7
	} else {
		valorFinal = valor * 1.0
	}
	return valorFinal
}
