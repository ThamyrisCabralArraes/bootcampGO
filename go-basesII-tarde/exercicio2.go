package main

import "fmt"

type loja struct {
	lista []produto
}

type produto struct {
	nome  string
	tipo  string
	preco float64
}

type Produto interface {
	CalcularCusto() float64
}

type Ecommerce interface {
	Adicionar(produto)
	Total() float64
}

func exercicio2() {
	store := novaLoja()
	store.Adicionar(novoProduto("Notebook", "medio", 1000.00))
	store.Adicionar(novoProduto("Smartphone", "grande", 2000.00))
	store.Adicionar(novoProduto("Tablet", "pequeno", 500.00))

	fmt.Println(store.Total())
}

func novoProduto(n, t string, p float64) produto {
	return produto{n, t, p}
}

func (p produto) CalcularCusto() float64 {
	switch p.tipo {
	case "medio":
		return p.preco * 0.03
	case "grande":
		return p.preco*0.06 + 2500
	default:
		return 0.00
	}
}

func (l *loja) Adicionar(p produto) {
	l.lista = append(l.lista, p)
}

func (l *loja) Total() float64 {
	var total float64
	for _, p := range l.lista {
		total += p.CalcularCusto() + p.preco
	}
	return total
}

func novaLoja() Ecommerce {
	return &loja{}
}

func CalcularCusto(p Produto) float64 {
	return p.CalcularCusto()
}
