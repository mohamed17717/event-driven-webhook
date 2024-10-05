package utils

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
