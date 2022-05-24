package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/guhkun13/go-pos/models"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	// import "gorm.io/driver/sqlite"
	// ref: https://gorm.io/docs/connecting_to_the_database.html#SQLite
	db, err := gorm.Open(sqlite.Open("api.sqlite"), &gorm.Config{})	
	if err != nil {
		log.Fatal("Failed to connect to the db \n", err.Error())
		os.Exit(2)
	}
	
	log.Println("Database connected!")

	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running migrations")
	// TODO : Add migrations
	db.AutoMigrate(&models.User{}, &models.Product{})

	Database = DbInstance{Db: db}
}