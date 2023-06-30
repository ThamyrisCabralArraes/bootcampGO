package products

type Service interface {
	GetAll() ([]Product, error)
	Store(nome, tipo string, quantidade int, preco float64) (Product, error)
	Update(id int, nome, tipo string, quantidade int, preco float64) (Product, error)
	UpdateName(id int, nome string) (Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}

}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(nome, tipo string, quantidade int, preco float64) (Product, error) {
	lastID, err := s.repository.lastID()
	if err != nil {
		return Product{}, err
	}
	lastID++

	product, err := s.repository.Store(lastID, nome, tipo, quantidade, preco)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Update(id int, nome, tipo string, quantidade int, preco float64) (Product, error) {
	return s.repository.Update(id, nome, tipo, quantidade, preco)
}

func (s *service) UpdateName(id int, nome string) (Product, error) {
	return s.repository.UpdateName(id, nome)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
