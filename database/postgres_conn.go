package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/guhkun13/go-pos/models"
	"github.com/guhkun13/go-pos/internal"
)

var DBPgsql *gorm.DB

func PostgresConnection() {
	var err error
	// import "gorm.io/driver/postgres"
	// ref: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL

	dsn, _ := internal.ConnectionURLBuilder("postgres")
	log.Println(dsn)

	DBPgsql, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("Failed connect to database")
	}
	
	log.Println("Connection opened to database")
	
	// Migrate DB
	tables := []string{"User", "Product"}

	DBPgsql.AutoMigrate(&models.User{})
	DBPgsql.AutoMigrate(&models.Product{})
	
	log.Println("Database migrated : ", tables)
}