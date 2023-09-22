package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func CheckId() fiber.Handler {

	return func(c *fiber.Ctx) error {

		tokenString := extractToken(c)

		id := c.Params("id")

		token, err := jwt.Parse(tokenString, jwtKeyFunc)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Failing chicking the role", "data": err.Error()})
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["user_id"] != id {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "You can only edit your profile", "data": nil})
		}
		return c.Next()
	}
}
