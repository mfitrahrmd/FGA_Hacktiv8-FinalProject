package domain

import (
	"context"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/infrastructure/infra_crypto"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Age      int     `json:"age"`
	Photos   []Photo `json:"photos"`
}

type UserRepository interface {
	FindOneById(ctx context.Context, id *uint) (*User, error)
	FindOneByEmail(ctx context.Context, email *string) (*User, error)
	FindOneByUsername(ctx context.Context, username *string) (*User, error)
	InsertOne(ctx context.Context, user *User) (*User, error)
	UpdateOneById(ctx context.Context, id *uint, user *User) (*User, error)
	DeleteOneById(ctx context.Context, id *uint) (*User, error)
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// hash user password
	hashedPassword, err := infra_crypto.HashPassword(u.Password)
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

	hashedPassword, err := infra_crypto.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	return nil
}
