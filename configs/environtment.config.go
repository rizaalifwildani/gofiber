package configs

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnvirontment() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. \n", err)
	}
}
