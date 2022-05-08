package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	user_model "github.com/guhkun13/go-pos/apps/user/model"
	product_model "github.com/guhkun13/go-pos/apps/product/model"
	"github.com/guhkun13/go-pos/pkg/utils"
)

var DB *gorm.DB

func PostgresConnection() {
	var err error
	// import "gorm.io/driver/postgres"
	// ref: https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL

	dsn, _ := utils.ConnectionURLBuilder("postgres")
	log.Println(dsn)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("Failed connect to database")
	}
	
	log.Println("Connection opened to database")
	
	// Migrate DB
	tables := []string{"User", "Product"}

	DB.AutoMigrate(&user_model.User{})
	DB.AutoMigrate(&product_model.Product{})
	
	log.Println("Database migrated : ", tables)
}