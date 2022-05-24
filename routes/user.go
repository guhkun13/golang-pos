package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/guhkun13/go-pos/database"
	"github.com/guhkun13/go-pos/internal"
	"github.com/guhkun13/go-pos/models"
)

func GetUsers(c *fiber.Ctx) error {
	conn := database.Database.Db
	var users []models.User
	
	conn.Find(&users)
	
	if len(users) == 0 {
		ret := internal.ReturnError(internal.DataNotFound, nil)
		log.Println(ret)
		
		return c.Status(404).JSON(ret)
	}
	
	ret := internal.ReturnOK(internal.OK, users)
	return c.Status(200).JSON(ret)
}

func CreateUser(c *fiber.Ctx) error {
	log.Printf("@func: CreateUser")
	
	// get db connection
	conn := database.DbSql
	
	// init user instance
	user := new(models.User)
	
	c.BodyParser(&user)
	// validate input first
	errors := internal.ValidateInput(&user)
	
	if errors != nil {
		return c.JSON(errors)		
	}
	
	// try to parse input to user instance
	if err := c.BodyParser(&user); err != nil {
		log.Println("err at body parser", err)
		ret := internal.ReturnError(internal.ReviewInput, err)
		
		return c.Status(500).JSON(ret)
	}
	
	// generate new ID
	user.ID = uuid.New()
	// log.Println("user to be created: ", user)
	
	// try to save to db
	if err := conn.Create(&user).Error; err != nil {
		ret := internal.ReturnError(internal.CannotCreateData, err)
		
		log.Println("err save to db: ", ret)
		return c.Status(500).JSON(ret)
	}
	ret := internal.ReturnOK(internal.OK, user)
	
	return c.Status(200).JSON(ret)
}

func GetUser(c *fiber.Ctx) error {
	// establish db conn
	conn := database.DbSql
	
	// get id
	id := c.Params("id")
	
	// prepare var 
	var user models.User
	
	// try get user
	if err := conn.Find(&user, "id = ?", id).Error; err != nil {
		log.Println("error find data: ", id)
		ret := internal.ReturnError(internal.DataNotFound, nil)
		
		return c.Status(404).JSON(ret)
	}
	
	// ok get user
	ret := internal.ReturnOK(internal.OK, user)
	
	return c.Status(200).JSON(ret)
}

func DeleteUser(c *fiber.Ctx) error {
	// establish db conn
	conn := database.DbSql
	
	// get id
	id := c.Params("id")
	
	// prepare var 
	var user models.User
	
	// try delete user
	if err := conn.Find(&user, "id = ?", id).Error; err != nil {
		log.Println("error find data: ", id)
		ret := internal.ReturnError(internal.DataNotFound, err)
		
		return c.Status(404).JSON(ret)
	}

	if user.ID == uuid.Nil {
		log.Println("user ID is nil", user.ID)

		ret := internal.ReturnError(internal.DataNotFound, nil)
		
		return c.Status(404).JSON(ret)
	}
	
	// get user berhasil, try to delete it 
	if err := conn.Delete(&user, "id = ?", id); err != nil {
		log.Println("gagal delete")
		ret := internal.ReturnError("Failed delete data", err)

		return c.Status(500).JSON(ret)
	}

	// ok get user
	ret := internal.ReturnOK(internal.OK, user)
	
	return c.Status(200).JSON(ret)
}