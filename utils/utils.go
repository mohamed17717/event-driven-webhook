package utils

import (
	"fmt"
	"github.com/joho/godotenv"
)

func CheckErr(err error, isPanic bool) {
	if err != nil {
		if isPanic {
			panic(err)
		}
		fmt.Println(err)
	}
}

func LoadEnv() {
	// load .env file
	err := godotenv.Load(".env")
	CheckErr(err, false)
}
