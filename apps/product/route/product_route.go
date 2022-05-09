package product_route

import (
	"log"

	"github.com/gofiber/fiber/v2"

	product_controller "github.com/guhkun13/go-pos/apps/product/controller"
)


func SetupProductRoutes(router fiber.Router) {
	log.Println("SetupProductRoutes")

	rg := router.Group("product")

	rg.Get("/", product_controller.GetProducts)	
}