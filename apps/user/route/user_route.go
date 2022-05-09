package user_route

import (
	"log"

	"github.com/gofiber/fiber/v2"

	user_controller "github.com/guhkun13/go-pos/apps/user/controller"
)

func SetupUserRoutes(router fiber.Router) {
	log.Println("SetUserRoutes")

	rg := router.Group("user")

	rg.Get("/", user_controller.GetUsers)
	rg.Post("/", user_controller.CreateUser)
	rg.Get("/:id", user_controller.GetUser)
	rg.Delete("/:id", user_controller.DeleteUser)
}