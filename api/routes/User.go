package routes

import (
	"github.com/gofiber/fiber/v2"

	"clean_architecture/api/handlers"
	"clean_architecture/pkg/user"
)

func UserRoutes(app fiber.Router, service user.Service) {
	app.Get("/users", handlers.IndexHandler(service))
	app.Post("/user", handlers.CreateHandler(service))
	app.Get("/user/:id", handlers.ShowHandler(service))
	app.Put("/user/:id", handlers.UpdateHandler(service))
	app.Delete("/user/:id", handlers.DeleteHandler(service))
}
