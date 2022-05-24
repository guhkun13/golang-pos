package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/guhkun13/go-pos/database"
	"github.com/guhkun13/go-pos/models"
)

type RegisterUserDto struct {
	Name 	string		`json:"name" validate:"required"`
	Email string
	Password string	
}


type UserLoginDto struct {
	Email string
	Password string
}

const SecretKey = "ThisIsMySecretKey"

func Register(c *fiber.Ctx) error {
	var data RegisterUserDto

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 12)

	if data.Name == "" || data.Email == "" {
		return c.Status(http.StatusBadRequest).JSON("Fix your data payload")
	}

	user := models.User{
		ID: 			uuid.New(),
		Name:     data.Name,
		Email:    data.Email,
		Password: string(password),		
	}

	if err := database.DbSql.Create(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(user)
}

func Login(c *fiber.Ctx) error {

	log.Println("method login!")
	var account UserLoginDto
	// user := c.FormValue("user")
	// pass := c.FormValue("pass")

	if err := c.BodyParser(&account); err != nil {
		log.Println("Error on parsing" ,err.Error())
	}

	log.Println("success parsing body:", account)

	// Throws Unauthorized error
	// if account.u != "john" || pass != "doe" {
	// 	return c.SendStatus(fiber.StatusUnauthorized)
	// }

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}