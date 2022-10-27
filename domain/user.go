package domain

import (
	"context"

	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/helper/helper_crypto"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username    string       `json:"username,omitempty"`
	Email       string       `json:"email,omitempty"`
	Password    string       `json:"-"`
	Age         int          `json:"age,omitempty"`
	Photo       *Photo       `json:"photo,omitempty"`
	Comments    []Comment    `json:"comments,omitempty"`
	SocialMedia *SocialMedia `json:"social_media,omitempty"`
}

type UserRegister struct {
	Username string `json:"username" binding:"required,uniqueUsername"`
	Email    string `json:"email" binding:"required,email,uniqueEmail"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,min=8"`
}

type UserRegisterResponse struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdateData struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
}

type UserUpdateDataResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserDeleteResponse struct {
	Message string `json:"message"`
}

type UserRepository interface {
	FindAndCount(ctx context.Context, user *User) (*int, error)
	FindAll(ctx context.Context) (*[]User, error)
	FindOneById(ctx context.Context, id *uint) (*User, error)
	FindOneByEmail(ctx context.Context, email *string) (*User, error)
	FindOneByUsername(ctx context.Context, username *string) (*User, error)
	InsertOne(ctx context.Context, user *User) (*User, error)
	UpdateOneById(ctx context.Context, id *uint, user *User) (*User, error)
	DeleteOneById(ctx context.Context, id *uint) (*User, error)
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := helper_crypto.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if u.Password == "" {
		return nil
	}

	hashedPassword, err := helper_crypto.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}
