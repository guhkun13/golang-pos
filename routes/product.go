package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/guhkun13/go-pos/database"
	"github.com/guhkun13/go-pos/models"
)


func GetProducts(c *fiber.Ctx) error {
	log.Println("GetProducts")
	conn := database.DbSql

	var products []models.Product

	conn.Find(&products)

	log.Println("products found : ", products)

	if len(products) == 0 {
		ret := fiber.Map{"status": "error","message" : "no data was found","data": nil}
		log.Println(ret)
		
		return c.Status(404).JSON(ret)
	}

	// return c.JSON(products)
	ret := fiber.Map{"status": "success", "message" : "Data found", "data": nil}
	log.Println(ret)

	return c.JSON(ret)
}