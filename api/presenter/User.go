package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"clean_architecture/pkg/entities"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func UserSuccessResponse(data *entities.User) *fiber.Map {
	user := User{
		ID:       data.ID,
		Name:     data.Name,
		Username: data.Username,
		Password: data.Password,
	}

	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

func UsersSuccessResponse(data *[]User) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
