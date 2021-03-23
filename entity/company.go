package entity

type Company struct {
	ID    int
	Email string
	Name  string
}

func NewCompany(email, name string) (*Company, error) {
	company := &Company{
		Email: email,
		Name:  name,
		ID:    0,
	}

	return company, nil
}
