package products

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
