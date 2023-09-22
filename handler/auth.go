package handler

import (
	"Sviluppo/go/go-fiber/config"
	"Sviluppo/go/go-fiber/database"
	"Sviluppo/go/go-fiber/models"
	"errors"
	"net/mail"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*models.User, error) {

	db := database.DB.Db
	var user models.User

	if err := db.Where(&models.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil

}

func getUserByUsername(u string) (*models.User, error) {
	db := database.DB.Db
	var user models.User

	if err := db.Where(&models.User{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Register(c *fiber.Ctx) error {

	type RegisterInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	register := new(models.User)
	db := database.DB.Db

	if err := c.BodyParser(&register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON((fiber.Map{"status": "error", "message": "Error on register request", "data": err.Error()}))
	}

	if existName, _ := getUserByUsername(register.Username); existName.Username != "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "This Username already exists", "data": register.Username})
	}

	if existEmail, _ := getUserByEmail(register.Email); existEmail.Email != "" {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "This Email already exists", "data": register.Email})
	}

	if hashPass, err := hashPassword(register.Password); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "errors": err.Error()})
	} else {
		register.Password = hashPass
	}

	if err := db.Create(register).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "errors": err.Error()})
	}

	registerInput := RegisterInput{
		Username: register.Username,
		Email:    register.Email,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": registerInput})
}

func Login(c *fiber.Ctx) error {

	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}

	type UserData struct {
		ID       uuid.UUID `gorm:"type:uuid"`
		Username string    `json:"username"`
		Email    string    `json:"email"`
		Password string    `json:"password"`
		Role     string    `json:"role"`
	}

	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err.Error()})
	}

	identity := input.Identity
	pass := input.Password
	userModel, err := new(models.User), *new(error)

	if isEmail(identity) {
		userModel, err = getUserByEmail(identity)
	} else {
		userModel, err = getUserByUsername(identity)
	}

	if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	} else {
		userData = UserData{
			ID:       userModel.ID,
			Username: userModel.Username,
			Email:    userModel.Email,
			Password: userModel.Password,
			Role:     userModel.Role,
		}
	}

	if !CheckPasswordHash(pass, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.ID
	claims["user_role"] = userData.Role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "failed to sign JWT", "data": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})

}
