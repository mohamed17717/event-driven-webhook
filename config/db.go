package config

import (
	"errors"
	"event-driven-webhook/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	databaseUrl := os.Getenv("DATABASE_URL")
	utils.FailOnError(errors.New("missed env variables"), "DATABASE_URL not set")

	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	utils.FailOnError(err, "Cannot connect to database")

}
