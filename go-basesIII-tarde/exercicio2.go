package main

type user struct {
	Nome      string
	Sobrenome string
	Email     string
}

type Produtos struct {
	Produtos []produto
}

type produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func novoProduto(p ...produto) Produtos {
	var produtos []produto

	for _, item := range p {
		produtos = append(produtos, item)
	}

	pessoa := "Thamyris"
	adicionarProdutos(&pessoa, &produtos)
	return Produtos{Produtos: produtos}
}

func adicionarProdutos(n *string, p *[]produto) {

}

func exercicio2() {
	p1 := produto{Nome: "Produto 1", Preco: 10.99, Quantidade: 5}
	// p2 := produto{Nome: "Produto 2", Preco: 19.99, Quantidade: 3}
	// p3 := produto{Nome: "Produto 3", Preco: 7.5, Quantidade: 8}

	novoProduto(p1)

}
