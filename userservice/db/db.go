package db

import (
	"fmt"
	"log"
	"os"
	"user_service/schemas"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectToDb() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("postgres://%v:%v@localhost:5432/%v", user, password, db)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err := DB.AutoMigrate(&schemas.UserData{}); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
}
