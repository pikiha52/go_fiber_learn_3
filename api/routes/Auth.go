package routes

import (
	"github.com/gofiber/fiber/v2"

	"clean_architecture/api/handlers"
	"clean_architecture/pkg/auth"
)

func AuthRoutes(app fiber.Router, service auth.Service) {
	app.Post("/auth/signin", handlers.SigninHandler(service))
}
