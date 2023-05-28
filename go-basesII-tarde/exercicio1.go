package main

import "fmt"

func exercicio1() {
	type Aluno struct {
		Nome         string
		Sobrenome    string
		RG           int
		DataAdmissao int
	}

	aluno1 := Aluno{"Maria", "Silva", 1234445, 12032003}
	fmt.Println(aluno1)
}
