package middleware

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ValidateProdcut(c *fiber.Ctx) error {

	type ProductValidationStruct struct {
		ID          uuid.UUID `gorm:"type:uuid; validate: uuid"`
		Title       string    `json:"title" validate:"required, min=3"`
		Description string    `json:"description" validate:"required, min=10"`
		Amount      int       `json:"amount" validate:"required, number"`
	}

	var bodyReq ProductValidationStruct

	if err := c.BodyParser(&bodyReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Could not validate data", "data": err.Error()})
	}

	validate := validator.New()

	if err := validate.Struct(bodyReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Can not validate data", "data": err.Error()})
	}

	return c.Next()
}

func ValidateUser(c *fiber.Ctx) error {
	type userValidationStruc struct {
		ID       uuid.UUID `gorm:"type:uuid; validate: uuid"`
		Username string    `json:"username" validate:"required, gte=4"`
		Email    string    `json:"email" validate:"required, email"`
		Password string    `json:"password" validate:"required, alphanumunicode, gte=6"`
	}

	var bodyReq userValidationStruc

	if err := c.BodyParser(&bodyReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Can not validate user", "data": err.Error()})
	}
	return c.Next()
}
