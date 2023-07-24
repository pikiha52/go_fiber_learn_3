package handlers

import (
	"github.com/gofiber/fiber/v2"

	"clean_architecture/api/presenter"
	"clean_architecture/pkg/auth"

)

func SigninHandler(service auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody presenter.AuthRequest

		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(500)
			return c.JSON(presenter.ErrorResponse(err))
		}

		result, err := service.SigninService(&requestBody)
		if err != nil {
			c.Status(500)
			return c.JSON(presenter.ErrorResponse(err))
		}

		return c.JSON(presenter.SuccessResponse(result))
	}
}
