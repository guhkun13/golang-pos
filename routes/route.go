package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"	
)

func setProtectedRoute(app *fiber.App){
	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))
}


func SetupRoutes(app *fiber.App) {
	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Halo Dunia")
	})

	authRoute := app.Group("auth")
	apiRoute := app.Group("api")
	productRoute := apiRoute.Group("products")
	userRoute := apiRoute.Group("users")

	authRoute.Post("/login", Login)
	authRoute.Post("/register", Register)
	
	userRoute.Get("/", GetUsers)
	setProtectedRoute(app)

	// product
	productRoute.Get("/", GetProducts)
	
	// user
	userRoute.Post("/", CreateUser)
	userRoute.Get("/:id", GetUser)
	userRoute.Delete("/:id", DeleteUser)
}