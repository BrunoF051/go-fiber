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

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func CheckRole() fiber.Handler {

	return func(c *fiber.Ctx) error {

		onlyToken := extractToken(c)

		token, err := jwt.Parse(onlyToken, jwtKeyFunc)

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
