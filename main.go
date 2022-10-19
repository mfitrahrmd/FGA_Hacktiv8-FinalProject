package main

import (
	"context"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/infrastructure/infra_gorm/database"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/infrastructure/infra_gorm/repository"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/user"
	"log"
)

func main() {
	env.LoadEnvFile()

	conn := database.GetPostgresGorm()
	conn.Debug().AutoMigrate(&domain.Photo{}, &domain.User{})

	userRepo := repository.NewUserRepository(conn)
	userUsecase := user.NewUserUsecase(userRepo)

	ctx := context.Background()

	id := "user-dd47b544daa74661b7b1cbc9353eb545"

	_, err := userUsecase.Delete(ctx, &id)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
