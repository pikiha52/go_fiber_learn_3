package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"clean_architecture/api/presenter"
	"clean_architecture/config"
	"clean_architecture/pkg/entities"
)

type Repository interface {
	SigninProcess(auth *presenter.AuthRequest) (*entities.Response, error)
}

type repository struct {
	Database *gorm.DB
}

func NewRepo(database *gorm.DB) Repository {
	return &repository{
		Database: database,
	}
}

func (r *repository) SigninProcess(auth *presenter.AuthRequest) (*entities.Response, error) {
	var user *entities.User
	r.Database.Where("username = ?", auth.Username).First(&user)

	if user.ID == uuid.Nil {
		return nil, errors.New("Username or password wrong!")
	}

	if !CheckPasswordHash(auth.Password, user.Password) {
		return nil, errors.New("Username or password wrong!")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return nil, err
	}

	response := entities.Response{
		Username:    user.Username,
		AccessToken: t,
	}

	return &response, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
