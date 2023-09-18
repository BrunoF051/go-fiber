package handler

import (
	"Sviluppo/go/go-fiber/database"
	"Sviluppo/go/go-fiber/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(models.User)

	err := c.BodyParser(user)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	err = db.Create(&user).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	return c.Status(500).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})

}

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []models.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

func GetSingleUser(c *fiber.Ctx) error {

	db := database.DB.Db

	id := c.Params("id")

	var user models.User

	db.First(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

func UpdateUser(c *fiber.Ctx) error {

	type UpdateUser struct {
		Username string `json:"username"`
	}

	db := database.DB.Db
	id := c.Params("id")
	var user models.User

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	var updaUserData UpdateUser
	err := c.BodyParser(&updaUserData)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	user.Username = updaUserData.Username

	db.Save(&user)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})

}

func DeleteUserById(c *fiber.Ctx) error {
	db := database.DB.Db
	var user models.User
	id := c.Params("id")

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
