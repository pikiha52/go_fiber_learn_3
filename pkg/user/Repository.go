package user

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"clean_architecture/api/presenter"
	"clean_architecture/pkg/entities"

)

type Repository interface {
	QueryAll() (*[]presenter.User, error)
	CreateOne(user *entities.User) (*entities.User, error)
	ShowOne(id string) (*entities.User, error)
	UpdateOne(id string, user *entities.User) (*entities.User, error)
	DeleteOne(id string) error
}

type repository struct {
	Database *gorm.DB
}

func NewRepo(database *gorm.DB) Repository {
	return &repository{
		Database: database,
	}
}

func (r *repository) QueryAll() (*[]presenter.User, error) {
	var users []presenter.User
	r.Database.Find(&users)

	return &users, nil
}

func (r *repository) CreateOne(user *entities.User) (*entities.User, error) {
	user.ID = uuid.New()

	hash, err := HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hash

	errSave := r.Database.Create(&user).Error

	if errSave != nil {
		return nil, errSave
	}

	return user, nil
}

func (r *repository) ShowOne(id string) (*entities.User, error) {
	var user *entities.User
	err := r.Database.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) UpdateOne(id string, user *entities.User) (*entities.User, error) {
	var userModel entities.User

	r.Database.Find(&userModel, "id = ?", id)

	if userModel.ID == uuid.Nil {
		return nil, errors.New("User not found!")
	}

	userModel.Name = user.Name
	userModel.Username = user.Username
	userModel.Password = user.Password

	r.Database.Save(&userModel)

	return &userModel, nil
}

func (r *repository) DeleteOne(id string) error {
	var user *entities.User

	r.Database.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return errors.New("user is not found!")
	}

	err := r.Database.Delete(&user, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
