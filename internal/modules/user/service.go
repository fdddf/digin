package user

type Service struct {
	// Add dependencies like DB connections here
}

type User struct {
	ID    string
	Name  string
	Email string
}

func NewService() (*Service, error) {
	return &Service{}, nil
}

func (s *Service) GetByID(id string) (*User, error) {
	return &User{
		ID:    id,
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil
}
