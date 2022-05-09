package root_route

import (
	"log"

	"github.com/gofiber/fiber/v2"

	root_controller "github.com/guhkun13/go-pos/apps/root/controller"
)

func SetRootRoutes(app *fiber.App, router fiber.Router) {
	log.Println("SetPublicRoutes")

	router.Get("/", root_controller.Welcome)

	static_folder := "/apps/root/statics"
	static_path := "."+static_folder
	
	app.Static("/apa", static_path)
}