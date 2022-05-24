package internal

import (
	"github.com/gofiber/fiber/v2"
)

const (
	OK string = "OK"
	DataNotFound string = "Data not found"
	ReviewInput string = "Please review your input"
	CannotCreateData string = "Cannot create data"
)

func ReturnError(msg string, data any) fiber.Map {
	
	result := fiber.Map{
		"status": false,
		"message": msg,
		"data": data,
	}
	
	return result
}

func ReturnOK(msg string, data any) fiber.Map {
	
	result := fiber.Map{
		"status": true,
		"message": msg,
		"data": data,
	}
	
	return result
}