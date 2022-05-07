package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"

	"github.com/guhkun13/go-pos/apps/router"
	"github.com/guhkun13/go-pos/pkg/database"
	"github.com/guhkun13/go-pos/pkg/utils"
)

func main() {

	log.Println("Start server")
	
	// auto reload
	fastergoding.Run()

	// create new fiber app
	app := fiber.New()
	
	// connect database
	database.PostgresConnection()

	// setup routes
	router.SetupRoutes(app)

	// start server
	utils.StartServer(app)
}