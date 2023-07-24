package presenter

import (
	"github.com/gofiber/fiber/v2"

	"clean_architecture/pkg/entities"
)

type AuthRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

func SuccessResponse(data *entities.Response) *fiber.Map {
	response := entities.Response{
		Username:    data.Username,
		AccessToken: data.AccessToken,
	}

	return &fiber.Map{
		"status": true,
		"data":   response,
		"error":  nil,
	}
}

func ErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
