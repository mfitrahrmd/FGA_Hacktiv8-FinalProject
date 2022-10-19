package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

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

	PG_USER = os.Getenv("PG_USER")
	PG_PASSWORD = os.Getenv("PG_PASSWORD")
	PG_HOST = os.Getenv("PG_HOST")
	PG_PORT = os.Getenv("PG_PORT")
	PG_DB = os.Getenv("PG_DB")

	JWT_KEY = os.Getenv("JWT_KEY")
	JWT_DURATION, _ = strconv.Atoi(os.Getenv("JWT_DURATION"))
}
