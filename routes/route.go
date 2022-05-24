package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Halo Dunia")
	})

	apiRoute := app.Group("api")
	
	productRoute := apiRoute.Group("products")
	productRoute.Get("/", GetProducts)
	
	userRoute := apiRoute.Group("users")
	userRoute.Get("/", GetUsers)
	userRoute.Post("/", CreateUser)
	userRoute.Get("/:id", GetUser)
	userRoute.Delete("/:id", DeleteUser)
}