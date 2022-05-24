package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/guhkun13/go-pos/models"
)

var DbSql *gorm.DB

func ConnectDbSQL() {
	var err error
	// import "gorm.io/driver/sqlite"
	// ref: https://gorm.io/docs/connecting_to_the_database.html#SQLite
	dbName := "gopos.db"
	DbSql, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to connect to database ", err.Error())
		os.Exit(2)
		}	
	log.Println("Database connected")	
	
	DbSql.Logger = logger.Default.LogMode(logger.Info)
	
	log.Println("Running migrations")
	// Migrate DB
	DbSql.AutoMigrate(&models.User{}, &models.Product{})
}