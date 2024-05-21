package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"taskservice/protos"
)

var DB *gorm.DB

type Task struct {
	gorm.Model
	protos.Task
}

func ConnectToDb() {
	dsn := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err := DB.AutoMigrate(&Task{}); err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
}
