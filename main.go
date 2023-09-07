package main

import (
	"Sviluppo/go/go-fiber/router"

	"Sviluppo/go/go-fiber/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDb()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	router.SetUpRoutes(app)

	app.Listen(":3000")
}
