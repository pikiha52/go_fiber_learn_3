package routes

import (
	"github.com/gofiber/fiber/v2"

	"clean_architecture/api/handlers"
	"clean_architecture/api/middleware"
	"clean_architecture/pkg/user"

)

func UserRoutes(app fiber.Router, service user.Service) {
	app.Get("/users", middleware.Protected(), handlers.IndexHandler(service))
	app.Post("/user", middleware.Protected(), handlers.CreateHandler(service))
	app.Get("/user/:id", middleware.Protected(), handlers.ShowHandler(service))
	app.Put("/user/:id", middleware.Protected(), handlers.UpdateHandler(service))
	app.Delete("/user/:id", middleware.Protected(), handlers.DeleteHandler(service))
}
