package user

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(router fiber.Router) {
	log.Println("SetUserRoutes")

	route := router.Group("user")

	route.Get("/", GetUsers)
}