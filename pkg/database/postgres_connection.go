package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/guhkun13/go-pos/apps/user"
	"github.com/guhkun13/go-pos/pkg/utils"
)

func PostgresConnection() {
	// import "gorm.io/driver/postgres"
	// ref: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL
	dsn, _ := utils.ConnectionURLBuilder("postgres")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed connect to database")
	}

	log.Println("Connection opened to database")

	// Migrate DB
	db.AutoMigrate(&user.UserModel{})

	log.Println("Database migrated")

	
}