package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	// load .env file
	err := godotenv.Load(".env")
	FailOnError(err, "Error loading .env file")
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func LogOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s\n", msg, err)
	}
}
