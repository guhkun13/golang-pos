package routes

import (
	"log"
	"net/http"
	"time"
	// "time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	// "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/guhkun13/go-pos/database"
	"github.com/guhkun13/go-pos/models"
)

const SecretKey = "ThisIsMySecretKey"

type RegisterUserDto struct {
	Name 	string		`json:"name" validate:"required"`
	Email string
	Password string	
}


type UserLoginDto struct {
	Email string			
	Password string
}

func Register(c *fiber.Ctx) error {
	var data RegisterUserDto

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data.Name == "" || data.Email == "" || data.Password == "" {
		msg := fiber.Map{"status": false, "message":"Fix your data payload"}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 12)

	user := models.User{
		ID: 			uuid.New(),
		Name:     data.Name,
		Email:    data.Email,
		Password: string(password),		
	}

	if err := database.DbSql.Create(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON("successfully create register new user")
}

func Login(c *fiber.Ctx) error {

	log.Println("method login!")
	var data UserLoginDto

	if err := c.BodyParser(&data); err != nil {
		log.Println("Error on parsing" ,err.Error())
	}

	if data.Email == "" || data.Password == "" {
		msg := fiber.Map{"status": false, "message":"Fix your data payload"}
		return c.Status(http.StatusBadRequest).JSON(msg)
	}

	log.Println("success parsing body:", data)

	var user models.User
	if err := database.DbSql.Find(&user, "email = ?", data.Email).Error; err != nil {
		log.Println("Error finding data", err.Error())
	}
	
	// log.Println(uuid.)
	if user.ID == uuid.Nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	log.Println("data found = ", user)

	log.Printf("Compare password input vs db =  \n %s \n %s \n ", []byte(data.Password), []byte(user.Password))

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		log.Println("Error on comparing password : ", err.Error())
		c.Status(fiber.StatusUnauthorized)

		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  user.Name,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}