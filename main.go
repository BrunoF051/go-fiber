package main

import (
	"Sviluppo/go/go-fiber/router"
	"log"

	"Sviluppo/go/go-fiber/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	database.ConnectDb()

	app := fiber.New()
	app.Use(logger.New())
	app.Static("/", "./public")
	router.SetUpRoutes(app)

	app.Use(cors.New())

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3000"))
}
