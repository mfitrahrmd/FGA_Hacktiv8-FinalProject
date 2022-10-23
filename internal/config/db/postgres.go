package db

import (
	"fmt"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
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

func Seed() {
	if conn == nil {
		StartPostgresGorm()
	}

	conn.Debug().Model(&domain.User{}).Save(&domain.User{
		Username: "rama1",
		Email:    "rama1@gmail.com",
		Password: "rama1",
		Age:      22,
	})

	conn.Debug().Model(&domain.Photo{}).Save(&domain.Photo{
		Title:    "Profile Picture",
		Caption:  "vacation",
		PhotoUrl: "https://rama1",
		UserId:   1,
	})

	conn.Debug().Model(&domain.Comment{}).Save(&domain.Comment{
		Message: "very cool",
		UserId:  1,
		PhotoId: 1,
	})

	conn.Debug().Model(&domain.SocialMedia{}).Save(&domain.SocialMedia{
		Name:           "Facebook",
		SocialMediaUrl: "https://facebook.com/rama1",
		UserId:         1,
	})
}
