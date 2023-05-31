package main

import (
	"fmt"
	"os"
)

type cliente struct {
	Arquivo  *os.File
	Nome     string
	RG       int
	Telefone int
	Endereco string
	ID       int
}

func (c *cliente) Cadastro(client ...cliente) {
	var err error
	c.Arquivo, err = os.Open("customers.txt")

	defer func() {
		if err != nil {
			panic("erro: o arquivo indicado não foi encontrado ou está danificado")
		}
	}()

	clienteCadastrado := append([]cliente{}, client...)
	fmt.Println("resposta2", clienteCadastrado)
}

func exercicio2() {
	id := 1
	c := cliente{}
	c.Cadastro(cliente{Nome: "Thamyris", RG: 123456, Telefone: 33339999, Endereco: "Rua da hora", ID: id})
}
