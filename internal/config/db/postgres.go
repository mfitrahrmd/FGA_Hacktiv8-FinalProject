package db

import (
	"fmt"
	"log"

	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn *gorm.DB

func StartPostgresGorm() {
	var err error

	conn, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", env.PG_HOST, env.PG_USER, env.PG_PASSWORD, env.PG_DB, env.PG_PORT)))
	if err != nil {
		log.Fatalln("error connecting to database :", err.Error())
	}
}

func GetPostgresGorm() *gorm.DB {
	if conn == nil {
		StartPostgresGorm()
	}

	return conn
}
