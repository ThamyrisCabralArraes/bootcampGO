package products

import (
	"fmt"
)

type Product struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Tipo       string  `json:"tipo"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}

var ps []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nome, tipo string, quantidade int, preco float64) (Product, error)
	lastID() (int, error)
	Update(id int, nome, tipo string, quantidade int, preco float64) (Product, error)
	UpdateName(id int, nome string) (Product, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (r *repository) lastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nome, tipo string, quantidade int, preco float64) (Product, error) {
	p := Product{id, nome, tipo, quantidade, preco}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

func (r *repository) Update(id int, nome, tipo string, quantidade int, preco float64) (Product, error) {
	p := Product{Nome: nome, Tipo: tipo, Quantidade: quantidade, Preco: preco}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d nao encontrado", id)
	}
	return p, nil
}

func (r *repository) UpdateName(id int, nome string) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Nome = nome
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d nao encontrado", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("produto %d nao encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil
}
