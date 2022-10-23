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

type UserRepository interface {
	FindAll(ctx context.Context) (*[]User, error)
	FindOneById(ctx context.Context, id *uint) (*User, error)
	FindOneByEmail(ctx context.Context, email *string) (*User, error)
	FindOneByUsername(ctx context.Context, username *string) (*User, error)
	InsertOne(ctx context.Context, user *User) (*User, error)
	UpdateOneById(ctx context.Context, id *uint, user *User) (*User, error)
	DeleteOneById(ctx context.Context, id *uint) (*User, error)
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// hash user password
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
