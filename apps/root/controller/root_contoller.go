package root_controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
)


func Welcome(c *fiber.Ctx) error {
	log.Println("func Welcome")

	return c.SendString("Welcome to the app")
}
