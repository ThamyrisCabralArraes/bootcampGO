package main

import (
	"criando-db/cmd/internal/products"
	"criando-db/cmd/internal/products/models"
	"fmt"
)

func main() {
	db := products.NewRepo()

	product := models.Product{
		Name:  "Produto 1",
		Type:  "Tipo 1",
		Count: 10,
		Price: 10.5,
	}

	id, err := db.Store(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	products, err := db.GetAll()
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Println(product)
	}
}
