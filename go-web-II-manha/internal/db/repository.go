package db

import "fmt"

type Request struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Senha int    `json:"idade"`
}

var login []Request
var lastID int

type Repository interface {
	Store(id int, nome string, senha int) (Request, error)
	GetAll() ([]Request, error)
	lastID() (int, error)
	Update(id int, nome string, senha int) (Request, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Request, error) {
	return login, nil
}

func (r *repository) lastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nome string, senha int) (Request, error) {
	l := Request{id, nome, senha}
	login = append(login, l)
	lastID = l.ID
	return l, nil
}

func (s *repository) Update(id int, nome string, senha int) (Request, error) {
	l := Request{ID: id, Nome: nome, Senha: senha}
	updated := false
	for i := range login {
		if login[i].ID == id {
			l.ID = id
			login[i] = l
			updated = true
		}
	}
	if !updated {
		return Request{}, fmt.Errorf("produto %d nao encontrado", id)
	}
	return l, nil
}
