package auth

import (
	"clean_architecture/api/presenter"
	"clean_architecture/pkg/entities"
)

type Service interface {
	SigninService(auth *presenter.AuthRequest) (*entities.Response, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) SigninService(auth *presenter.AuthRequest) (*entities.Response, error) {
	return s.repository.SigninProcess(auth)
}
