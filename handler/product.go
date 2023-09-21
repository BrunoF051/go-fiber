package handler

import (
	"Sviluppo/go/go-fiber/database"
	"Sviluppo/go/go-fiber/models"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func getProductByTitle(t string) (*models.Product, error) {
	db := database.DB.Db
	var product models.Product

	if err := db.Where(&models.Product{Title: t}).Find(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

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

	if existTitle, _ := getUserByEmail(product.Title); existTitle.Email != "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "This Title already exists", "data": product.Title})
	}

	db.Create(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Created product", "data": product})
}

func UpdateProduct(c *fiber.Ctx) error {

	type UpdateProductStruct struct {
		Title       string `json:"title" validate:"min=3"`
		Description string `json:"description" validate:"min=5"`
		Amount      int    `json:"amount" validate:"number"`
		Price       int    `json:"price" validate:"number"`
	}

	db := database.DB.Db
	var product models.Product
	id := c.Params("id")

	db.Find(&product, "id = ?", id)

	if product.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Product not found", "data": nil})
	}

	var updateProduct UpdateProductStruct

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	product.Title = updateProduct.Title
	product.Description = updateProduct.Description
	product.Amount = updateProduct.Amount
	product.Price = updateProduct.Price

	db.Save(&product)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product Found", "data": product})

}

func DeleteProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB.Db
	var product models.Product

	db.First(&product, "id = ?", id)

	if product.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	db.Delete(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
