package domain

import (
	"context"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/infrastructure/infra_crypto"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/infrastructure/infra_uuid"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Age      int     `json:"age"`
	Photo    []Photo `json:"photos"`
}

type UserRepository interface {
	FindOneById(ctx context.Context, id *string) (*User, error)
	FindOneByEmail(ctx context.Context, email *string) (*User, error)
	FindOneByUsername(ctx context.Context, username *string) (*User, error)
	InsertOne(ctx context.Context, user *User) error
	UpdateOneById(ctx context.Context, id *string, user *User) error
	DeleteOneById(ctx context.Context, id *string) error
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	// generate user id
	u.Id = infra_uuid.GenerateUUID("user")

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
