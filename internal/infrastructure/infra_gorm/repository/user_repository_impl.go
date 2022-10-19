package repository

import (
	"context"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	gorm *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		gorm: db,
	}
}

func (u userRepository) FindOneById(ctx context.Context, id *string) (*domain.User, error) {
	var result *domain.User

	if err := u.gorm.Debug().Model(&domain.User{}).Where("id = ?", *id).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u userRepository) FindOneByEmail(ctx context.Context, email *string) (*domain.User, error) {
	var result *domain.User

	if err := u.gorm.Debug().Model(&domain.User{}).Where("email = ?", *email).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u userRepository) FindOneByUsername(ctx context.Context, username *string) (*domain.User, error) {
	var result *domain.User

	if err := u.gorm.Debug().Model(&domain.User{}).Where("username = ?", *username).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u userRepository) InsertOne(ctx context.Context, user *domain.User) error {
	if err := u.gorm.Debug().Model(&domain.User{}).Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (u userRepository) UpdateOneById(ctx context.Context, id *string, user *domain.User) error {
	if err := u.gorm.Debug().Model(&domain.User{}).Where("id = ?", *id).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (u userRepository) DeleteOneById(ctx context.Context, id *string) error {
	if err := u.gorm.Debug().Model(&domain.User{}).Where("id = ?", id).Delete(&domain.User{}).Error; err != nil {
		return err
	}

	return nil
}
