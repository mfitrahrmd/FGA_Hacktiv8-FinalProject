package user

import (
	"context"
	"errors"

	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/helper/helper_crypto"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/helper/helper_jwt"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service"
	"gorm.io/gorm"
)

type UserUsecase interface {
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	Login(ctx context.Context, user *domain.User) (string, error)
	UpdateData(ctx context.Context, userId *uint, newDataUser *domain.User) (*domain.User, error)
	Delete(ctx context.Context, userId *uint) (*domain.User, error)
}

type userUsecase struct {
	UserRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) UserUsecase {
	return &userUsecase{
		UserRepository: userRepository,
	}
}

func (u userUsecase) isUserUnique(ctx context.Context, user *domain.User, userId ...uint) (bool, error) {
	countUser, err := u.UserRepository.FindAndCount(ctx, &domain.User{
		Email:    user.Email,
		Username: user.Username,
	})
	if err != nil {
		return false, err
	}

	if *countUser > 0 {
		return false, nil
	}

	return true, nil
}

func (u userUsecase) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	if unique, _ := u.isUserUnique(ctx, user); !unique {
		return nil, service.USERNAME_OR_EMAIL_ALREADY_EXIST
	}

	registeredUser, err := u.UserRepository.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return registeredUser, nil
}

func (u userUsecase) Login(ctx context.Context, user *domain.User) (string, error) {
	foundUser, err := u.UserRepository.FindOneByEmail(ctx, &user.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", service.USER_DOES_NOT_EXIST
		}
		return "", err
	}

	// compare hashed password
	isMatch := helper_crypto.ComparePassword(user.Password, foundUser.Password)
	if !isMatch {
		return "", service.INVALID_PASSWORD
	}

	// generate token
	token, err := helper_jwt.GenerateToken(env.JWT_KEY, helper_jwt.TokenPayload{
		Id:       foundUser.Id,
		Email:    foundUser.Email,
		Username: foundUser.Username,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u userUsecase) UpdateData(ctx context.Context, userId *uint, newDataUser *domain.User) (*domain.User, error) {
	_, err := u.UserRepository.FindOneById(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.USER_DOES_NOT_EXIST
		}
		return nil, err
	}

	unique, err := u.isUserUnique(ctx, newDataUser, *userId)
	if err != nil {
		return nil, err
	}

	if !unique {
		return nil, service.USERNAME_OR_EMAIL_ALREADY_EXIST
	}
	updatedUser, err := u.UserRepository.UpdateOneById(ctx, userId, newDataUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u userUsecase) Delete(ctx context.Context, userId *uint) (*domain.User, error) {
	_, err := u.UserRepository.FindOneById(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.USER_DOES_NOT_EXIST
		}
		return nil, err
	}

	deletedUser, err := u.UserRepository.DeleteOneById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return deletedUser, nil
}
