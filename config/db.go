package config

import (
	"event-driven-webhook/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	utils.CheckErr(err, true)

	fmt.Println("Database connected")
}
