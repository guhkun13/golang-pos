package user_controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	user_model "github.com/guhkun13/go-pos/apps/user/model"
	"github.com/guhkun13/go-pos/pkg/database"
	"github.com/guhkun13/go-pos/pkg/utils"
)

func GetUsers(c *fiber.Ctx) error {
	conn := database.DB
	var users []user_model.User
	
	conn.Find(&users)
	
	if len(users) == 0 {
		ret := utils.ReturnError(utils.DataNotFound, nil)
		log.Println(ret)
		
		return c.Status(404).JSON(ret)
	}
	
	ret := utils.ReturnOK(utils.OK, users)
	return c.Status(200).JSON(ret)
}

func CreateUser(c *fiber.Ctx) error {
	log.Printf("@func: CreateUser")
	
	// get db connection
	conn := database.DB
	
	// init user instance
	user := new(user_model.User)
	
	c.BodyParser(&user)
	// validate input first
	errors := utils.ValidateStruct(*user, utils.UserModel)
	
	if errors != nil {
		return c.JSON(errors)		
	}
	
	// try to parse input to user instance
	if err := c.BodyParser(&user); err != nil {
		log.Println("err at body parser", err)
		ret := utils.ReturnError(utils.ReviewInput, err)
		
		return c.Status(500).JSON(ret)
	}
	
	// generate new ID
	user.ID = uuid.New()
	// log.Println("user to be created: ", user)
	
	// try to save to db
	if err := conn.Create(&user).Error; err != nil {
		ret := utils.ReturnError(utils.CannotCreateData, err)
		
		log.Println("err save to db: ", ret)
		return c.Status(500).JSON(ret)
	}
	ret := utils.ReturnOK(utils.OK, user)
	
	return c.Status(200).JSON(ret)
}

func GetUser(c *fiber.Ctx) error {
	// establish db conn
	conn := database.DB
	
	// get id
	id := c.Params("id")
	
	// prepare var 
	var user user_model.User
	
	// try get user
	if err := conn.Find(&user, "id = ?", id).Error; err != nil {
		log.Println("error find data: ", id)
		ret := utils.ReturnError(utils.DataNotFound, nil)
		
		return c.Status(404).JSON(ret)
	}
	
	// ok get user
	ret := utils.ReturnOK(utils.OK, user)
	
	return c.Status(200).JSON(ret)
}