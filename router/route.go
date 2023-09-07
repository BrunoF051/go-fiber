package router

import (
	"Sviluppo/go/go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	v1.Get("/", handler.GetAllUsers)
	v1.Get("/:id", handler.GetSingleUser)
	v1.Post("/", handler.CreateUser)
	v1.Put(":id", handler.UpdateUser)
	v1.Delete("/:id", handler.DeleteUserById)
}
