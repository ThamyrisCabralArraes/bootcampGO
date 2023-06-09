package db

type Servide interface {
	Store(nome string, senha int) (Request, error)
	GetAll() ([]Request, error)
	Update(id int, nome string, senha int) (Request, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Servide {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Request, error) {
	lg, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return lg, nil
}

func (s *service) Store(nome string, senha int) (Request, error) {
	lastID, err := s.repository.lastID()
	if err != nil {
		return Request{}, err
	}

	lastID++

	login, err := s.repository.Store(lastID, nome, senha)
	if err != nil {
		return Request{}, err
	}
	return login, nil
}

func (s *service) Update(id int, nome string, senha int) (Request, error) {
	return s.repository.Update(id, nome, senha)
}
