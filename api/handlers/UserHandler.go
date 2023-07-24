package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"clean_architecture/api/presenter"
	"clean_architecture/pkg/entities"
	"clean_architecture/pkg/user"

)

func IndexHandler(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.QueryAllService()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}
		return c.JSON(presenter.UsersSuccessResponse(fetched))
	}
}

func CreateHandler(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.User

		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		validate := validator.New()
		errValidate := validate.Struct(requestBody)

		if errValidate != nil {
			c.Status(400)
			return c.JSON(presenter.UserErrorResponse(errValidate))
		}

		result, err := service.InsertOneService(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

func ShowHandler(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		result, err := service.ShowOneService(id)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

func UpdateHandler(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		var requestBody entities.User

		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		result, err := service.UpdateOneService(id, &requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(presenter.UserSuccessResponse(result))
	}
}

func DeleteHandler(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := service.DeleteOneService(id)

		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.UserErrorResponse(err))
		}

		return c.JSON(fiber.Map{"status": true})
	}
}
