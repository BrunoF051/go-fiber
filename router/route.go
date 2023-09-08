package router

import (
	"Sviluppo/go/go-fiber/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("login", handler.Login)

	// User
	user := api.Group("/user")
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", handler.CreateUser)
	user.Put(":id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUserById)
}
