package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var HTTP_PORT int

func InitPort() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	HTTP_PORT, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		panic(err)
	}
}
