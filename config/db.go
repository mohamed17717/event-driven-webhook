package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		fmt.Println("DATABASE_URL not set")
		return
	}

	var err error
	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to database")
	}

	fmt.Println("Database connected")
}
