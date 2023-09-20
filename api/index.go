package handler

import (
	"Sviluppo/go/go-fiber/database"
	"Sviluppo/go/go-fiber/router"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)
}

func handler() http.HandlerFunc {

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

	return adaptor.FiberApp(app)
}
