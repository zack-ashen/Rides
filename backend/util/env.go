package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func AccessToken() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("%v", err)
	}

	return os.Getenv("ACCESS_TOKEN_SECRET")
}
