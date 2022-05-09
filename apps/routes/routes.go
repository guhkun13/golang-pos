package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	root_route "github.com/guhkun13/go-pos/apps/root/route"
	user_route "github.com/guhkun13/go-pos/apps/user/route"
	product_route "github.com/guhkun13/go-pos/apps/product/route"
)

func SetupRoutes(app *fiber.App) {
	const (
		root_path = "/"
	)

	app.Static("/static", "./apps/statics")

	// root
	root := app.Group(root_path, logger.New())

	// set app route
	root_route.SetRootRoutes(app, root)
	
	// API
	Setup_Api(app)
}


func Setup_Api(app *fiber.App) {
	const (
		path string = "/api"
	)
	api := app.Group(path, logger.New())

	// add app route
	user_route.SetupUserRoutes(api)
	product_route.SetupProductRoutes(api)
}