package main

import "fmt"

type error interface {
	Error() string
}

type myError struct {
	status string
	msg    string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%v - %v", e.status, e.msg)
}

func exercicio1(s float64) (string, error) {
	if s < 15.000 {
		return " ", &myError{
			status: "",
			msg:    "erro: O salário digitado não está dentro do valor mínimo",
		}
	}
	return "Necessário pagamento de imposto.", nil
}
