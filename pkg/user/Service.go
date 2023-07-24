package user

import (
	"clean_architecture/api/presenter"
	"clean_architecture/pkg/entities"
)

type Service interface {
	QueryAllService() (*[]presenter.User, error)
	InsertOneService(user *entities.User) (*entities.User, error)
	ShowOneService(id string) (*entities.User, error)
	UpdateOneService(id string, user *entities.User) (*entities.User, error)
	DeleteOneService(id string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) QueryAllService() (*[]presenter.User, error) {
	return s.repository.QueryAll()
}

func (s *service) InsertOneService(user *entities.User) (*entities.User, error) {
	return s.repository.CreateOne(user)
}

func (s *service) ShowOneService(id string) (*entities.User, error) {
	return s.repository.ShowOne(id)
}

func (s *service) UpdateOneService(id string, user *entities.User) (*entities.User, error) {
	return s.repository.UpdateOne(id, user)
}

func (s *service) DeleteOneService(id string) error {
	return s.repository.DeleteOne(id)
}
