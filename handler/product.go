package handler

import (
	"Sviluppo/go/go-fiber/database"
	"Sviluppo/go/go-fiber/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllProducts(c *fiber.Ctx) error {
	db := database.DB.Db
	var products []models.Product

	db.Find(&products)

	return c.JSON(fiber.Map{"status": "success", "message": "All products", "data": products})
}

func GetProducById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB.Db
	var product models.Product

	db.Find(&product, id)

	if product.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": product})

}

func CreateProduct(c *fiber.Ctx) error {
	db := database.DB.Db
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}

	db.Create(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

func DeleteProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB.Db
	var product models.Product

	db.First(&product, id)

	if product.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	db.Delete(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
