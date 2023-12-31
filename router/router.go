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
	auth.Post("/login", handler.Login)
	auth.Post("/register", middleware.ValidateUser(), handler.Register)

	// User
	user := api.Group("/user")
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetSingleUser)
	user.Post("/", middleware.Protected(), middleware.ValidateUser(), handler.CreateUser)
	user.Patch(":id", middleware.Protected(), middleware.ValidateUser(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUserById)

	// Product
	product := api.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", handler.GetProducById)
	product.Post("/", middleware.Protected(), middleware.ValidateProdcut(), handler.CreateProduct)
	product.Patch("/:id", middleware.Protected(), middleware.ValidateProdcut(), handler.UpdateProduct)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProductByID)

}
