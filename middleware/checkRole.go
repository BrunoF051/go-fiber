package middleware

import (
	"Sviluppo/go/go-fiber/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return config.Config(("SECRET")), nil
}

func CheckRole() fiber.Handler {

	return func(c *fiber.Ctx) error {

		bearToken := c.Get("Authorization")

		onlyToken := strings.Split(bearToken, "")

		if len(onlyToken) == 2 {
			bearToken = onlyToken[1]
		} else if len(onlyToken) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Auth token not found", "data": nil})
		}

		token, err := jwt.Parse(bearToken, jwtKeyFunc)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Failing chicking the role", "data": nil})
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "admin" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Only users with 'admin' role can do that", "data": nil})
		}
		return c.Next()
	}
}
