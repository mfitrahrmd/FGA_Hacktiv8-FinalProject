package main

import (
	"context"
	"fmt"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/infrastructure/infra_gorm/database"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/infrastructure/infra_gorm/repository"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/photo"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/user"
	"log"
)

func main() {
	env.LoadEnvFile()

	conn := database.GetPostgresGorm()
	conn.Debug().AutoMigrate(&domain.User{}, &domain.Photo{})

	ur := repository.NewUserRepository(conn)
	_ = user.NewUserUsecase(ur)
	pr := repository.NewPhotoRepository(conn)
	pu := photo.NewPhotoUsecase(pr)

	ctx := context.WithValue(context.Background(), "userTokenPayload", user.TokenPayload{
		Id:       2,
		Email:    "rama2@gmail.com",
		Username: "rama2",
	})

	var photoId uint = 1

	u, err := pu.DeleteUserPhoto(ctx, &photoId)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(u)
}
