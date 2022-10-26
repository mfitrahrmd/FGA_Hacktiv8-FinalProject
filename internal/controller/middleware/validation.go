package middleware

import (
	"context"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
)

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}

func RegisterCustomValidation(ctx context.Context, userRepository domain.UserRepository) {
	v := binding.Validator.Engine().(*validator.Validate)
	v.RegisterValidation("uniqueEmail", func(fl validator.FieldLevel) bool {
		userEmail := fl.Field().String()

		countUser, _ := userRepository.FindAndCount(ctx, &domain.User{
			Email: userEmail,
		})

		if *countUser > 0 {
			return false
		}

		return true
	})
	v.RegisterValidation("uniqueUsername", func(fl validator.FieldLevel) bool {
		userUsername := fl.Field().String()

		countUser, _ := userRepository.FindAndCount(ctx, &domain.User{
			Username: userUsername,
		})

		if *countUser > 0 {
			return false
		}

		return true
	})
}
