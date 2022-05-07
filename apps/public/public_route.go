package public

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetPublicRoutes(router fiber.Router) {
	log.Println("SetPublicRoutes")

	router.Get("/", Welcome)
}