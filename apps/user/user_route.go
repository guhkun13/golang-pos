package user

import (
	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(router fiber.Router) {
	userRouter := router.Group("user")

	userRouter.Get("/", GetUsers)
}