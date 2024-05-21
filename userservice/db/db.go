package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"userservice/schemas"
)

var DB *gorm.DB

func ConnectToDb() {
	dsn := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err := DB.AutoMigrate(&schemas.UserData{}); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
}
