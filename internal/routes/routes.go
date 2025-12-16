package routes

import (
	"github.com/gofiber/fiber/v2"

	uh "github.com/ghv061101/RestApiAge/internal/handler"
)

func Register(app *fiber.App, h *uh.UserHandler) {
	api := app.Group("/api")
	api.Post("/users", h.CreateUser)
	api.Get("/users", h.ListUsers)
	api.Get("/users/:id", h.GetUser)
	api.Put("/users/:id", h.UpdateUser)
	api.Delete("/users/:id", h.DeleteUser)
}
