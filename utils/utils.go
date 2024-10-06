package utils

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}
