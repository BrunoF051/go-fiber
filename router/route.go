package router

import (
	"Sviluppo/go/go-fiber/handler"
	"Sviluppo/go/go-fiber/middleware"

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
	user.Patch(":id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUserById)

	// Product
	product := api.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", handler.GetProducById)
	product.Post("/", middleware.Protected(), handler.CreateProduct)
	product.Patch("/:id", middleware.Protected(), handler.UpdateProduct)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProductByID)

}
