package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/qinains/fastergoding"

	"github.com/guhkun13/go-pos/database"
	"github.com/guhkun13/go-pos/routes"
)

func main() {
	log.Println("Start server")
	
	// auto reload
	fastergoding.Run()

	// create new fiber app
	app := fiber.New()

	app.Use(logger.New())
	
	// connect database
	database.ConnectDbSQL()

	// setup routes
	routes.SetupRoutes(app)

	// start server
	// utils.StartServer(app)
	port := ":5000"
	log.Fatal(app.Listen(port))
	
}