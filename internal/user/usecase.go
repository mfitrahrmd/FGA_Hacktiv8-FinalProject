package user

import (
	"context"
	"errors"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/infrastructure/infra_crypto"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/infrastructure/infra_jwt"
	"gorm.io/gorm"
)

type TokenPayload struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type userUsecase struct {
	UserRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) *userUsecase {
	return &userUsecase{
		UserRepository: userRepository,
	}
}

func (u userUsecase) isUserUnique(ctx context.Context, user *domain.User, userId ...string) (bool, error) {
	foundwithEmail, errwithEmail := u.UserRepository.FindOneByEmail(ctx, &user.Email)
	if errwithEmail != nil && !errors.Is(errwithEmail, gorm.ErrRecordNotFound) {
		return false, errwithEmail
	}

	foundwithUsername, errwithUsername := u.UserRepository.FindOneByUsername(ctx, &user.Username)
	if errwithUsername != nil && !errors.Is(errwithUsername, gorm.ErrRecordNotFound) {
		return false, errwithUsername
	}

	if len(userId) > 0 {
		if foundwithEmail != nil && foundwithUsername != nil {
			if userId[0] == foundwithEmail.Id && userId[0] == foundwithUsername.Id {
				return true, nil
			}

			return false, nil
		}
	}

	return foundwithEmail == nil && foundwithUsername == nil, nil
}

func (u userUsecase) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	if unique, _ := u.isUserUnique(ctx, user); !unique {
		return nil, NewUserError(EMAIL_OR_USERNAME_ALREADY_EXIST)
	}

	err := u.UserRepository.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	registeredUser, err := u.UserRepository.FindOneByEmail(ctx, &user.Email)
	if err != nil {
		return nil, err
	}

	return registeredUser, nil
}

func (u userUsecase) Login(ctx context.Context, user *domain.User) (string, error) {
	foundUser, err := u.UserRepository.FindOneByEmail(ctx, &user.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", NewUserError(USER_NOT_FOUND)
		}
		return "", err
	}

	// compare hashed password
	isMatch := infra_crypto.ComparePassword(user.Password, foundUser.Password)
	if !isMatch {
		return "", NewUserError(INVALID_PASSWORD)
	}

	// generate token
	token, err := infra_jwt.GenerateToken(env.JWT_KEY, TokenPayload{foundUser.Id, foundUser.Email})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u userUsecase) UpdateData(ctx context.Context, userId *string, newDataUser *domain.User) (*domain.User, error) {
	_, err := u.UserRepository.FindOneById(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewUserError(USER_NOT_FOUND)
		}
		return nil, err
	}

	unique, err := u.isUserUnique(ctx, newDataUser, *userId)
	if err != nil {
		return nil, err
	}

	if !unique {
		return nil, NewUserError(EMAIL_OR_USERNAME_ALREADY_EXIST)
	}

	err = u.UserRepository.UpdateOneById(ctx, userId, newDataUser)
	if err != nil {
		return nil, err
	}

	updatedUser, err := u.UserRepository.FindOneById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u userUsecase) Delete(ctx context.Context, userId *string) (*domain.User, error) {
	foundUser, err := u.UserRepository.FindOneById(ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewUserError(USER_NOT_FOUND)
		}
		return nil, err
	}

	err = u.UserRepository.DeleteOneById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}
