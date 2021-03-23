package company

import (
	"github.com/marcelocpinheiro/communication-engine/entity"
	"github.com/marcelocpinheiro/communication-engine/infrastructure/repository"
)

type Service struct {
	repo *repository.CompanyMySQL
}

func NewService(r *repository.CompanyMySQL) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateCompany(name, email string) (int64, error) {
	e, err := entity.NewCompany(email, name)
	if err != nil {
		return -1, nil
	}
	return s.repo.Create(e)
}
