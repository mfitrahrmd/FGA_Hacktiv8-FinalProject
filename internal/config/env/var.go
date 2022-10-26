package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// server variables
var HOST string = "localhost"
var PORT string = "8001"

// postgres variables
var PG_USER string
var PG_PASSWORD string
var PG_HOST string
var PG_PORT string
var PG_DB string

// jwt variables
var JWT_KEY string
var JWT_DURATION int

func LoadEnvFile() {
	wd, _ := os.Getwd()
	err := godotenv.Load(wd + "/.env")
	if err != nil {
		log.Fatalln(err.Error())
	}

	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")

	PG_USER = os.Getenv("PG_USER")
	PG_PASSWORD = os.Getenv("PG_PASSWORD")
	PG_HOST = os.Getenv("PG_HOST")
	PG_PORT = os.Getenv("PG_PORT")
	PG_DB = os.Getenv("PG_DB")

	JWT_KEY = os.Getenv("JWT_KEY")
	JWT_DURATION, _ = strconv.Atoi(os.Getenv("JWT_DURATION"))
}
