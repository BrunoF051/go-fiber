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
	app.Use(cors.New())

	router.SetUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
