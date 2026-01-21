package main

import (
	"cashinvoice-assignment/internal/config"
	"cashinvoice-assignment/internal/database"
	"cashinvoice-assignment/internal/handler"
	"cashinvoice-assignment/internal/model"
	"cashinvoice-assignment/internal/repository"
	"cashinvoice-assignment/internal/router"
	"cashinvoice-assignment/internal/service"
	"cashinvoice-assignment/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := config.Load()

	database.Connect(cfg)

	database.DB.AutoMigrate(
		&model.User{},
		&model.Todo{},
	)
	app := fiber.New()
	worker := utils.NewAutoCompleteWorker(cfg.AutoCompleteDelayMinutes)
	worker.Start(5)
	defer worker.Stop()
	userRepository := repository.NewUserRepository(database.DB)
	todoRepository := repository.NewTodoRepository(database.DB)
	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)
	todoService := service.NewTodoService(todoRepository, worker)
	todoHandler := handler.NewTodoHandler(todoService)
	router.Setup(app, authHandler, todoHandler)

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
	app.Listen(":8080")
}
