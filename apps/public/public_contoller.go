package public

import (
	"log"

	"github.com/gofiber/fiber/v2"
)


func Welcome(c *fiber.Ctx) error {
	log.Printf("call func Welcome")

	return c.SendString("Welcome")
}