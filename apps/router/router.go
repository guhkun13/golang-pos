package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/guhkun13/go-pos/apps/user"
	"github.com/guhkun13/go-pos/apps/public"
)

func SetupRoutes(app *fiber.App) {
	// create group based on app or path
	
	// root
	root := app.Group("/", logger.New())
	public.SetPublicRoutes(root)
	
	// auth
	// auth := app.Get("/auth", logger.New())
	
	// api
	api := app.Group("/api", logger.New())
	user.SetUserRoutes(api)
}