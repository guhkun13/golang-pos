package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	log.Println("Start Server")
	// Build Fiber connection URL.
	fiberConnURL, _ := ConnectionURLBuilder("fiber")
	log.Println("fiberConnURL: ", fiberConnURL)


	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
