package main

import (
	"errors"
	"fmt"
	"os"
)

type produto struct {
	id         int
	quantidade int
	preco      float64
}

func (p produto) guardarArquivos() string {
	return fmt.Sprintf("%d, %d, %.2f\n", p.id, p.quantidade, p.preco)
}

func (p produto) guardarArquivos2() string {
	return "id, quantidade, preco\n"
}

func lerArquivos(caminho string, produtos []produto) error {
	if len(produtos) == 0 {
		return errors.New("quantidade invalida")
	}

	file, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("Erro ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	p := produtos[0]

	if _, err = file.WriteString(p.guardarArquivos2()); err != nil {
		return fmt.Errorf("erro ao gerar cabeçalho: %w", err)
	}

	for _, p = range produtos {
		if _, err = file.WriteString(p.guardarArquivos()); err != nil {
			return fmt.Errorf("erro ao salvar linha: %w", err)
		}
	}
	return nil
}

func main() {
	produtos := []produto{
		{
			id:         1,
			quantidade: 2,
			preco:      2.90,
		},
		{
			id:         2,
			quantidade: 2,
			preco:      2.90,
		},
		{
			id:         3,
			quantidade: 2,
			preco:      2.90,
		},
	}
	lerArquivos("produtos", produtos)
}
