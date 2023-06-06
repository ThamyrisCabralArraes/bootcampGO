package products

type Service interface {
	GetAll() ([]Product, error)
	Store(nome, tipo string, quantidade int, preco float64) (Product, error)
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
