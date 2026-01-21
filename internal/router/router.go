package router

import (
	"cashinvoice-assignment/internal/handler"
	"cashinvoice-assignment/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, authHandler *handler.AuthHandler, todoHandler *handler.TodoHandler) {
	auth := app.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	protected := app.Group("/api", middleware.AuthMiddleware())

	protected.Post("/todos", todoHandler.Create)
	protected.Get("/todos", todoHandler.List)
	protected.Put("/todos/:id", todoHandler.Update)
	protected.Delete("/todos/:id", todoHandler.Delete)

	protected.Get("/profile", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id": c.Locals("user_id"),
			"email":   c.Locals("email"),
		})
	})
}
