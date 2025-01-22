package product

type Service struct {
	// Add dependencies here
}

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewService() (*Service, error) {
	return &Service{}, nil
}

func (s *Service) GetProduct(id string) (*Product, error) {
	return &Product{
		ID:    id,
		Name:  "Example Product",
		Price: 29.99,
	}, nil
}
