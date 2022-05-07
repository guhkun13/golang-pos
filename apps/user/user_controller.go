package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	log.Printf("call func GetUsers")
	fmt.Printf("c = %s ", c)

	return c.Status(http.StatusOK).SendString("List Users")
}