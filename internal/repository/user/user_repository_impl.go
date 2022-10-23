package user

import (
	"context"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	gorm *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		gorm: db,
	}
}

func (u userRepository) FindAll(ctx context.Context) (*[]domain.User, error) {
	var result []domain.User

	if err := u.gorm.Model(&domain.User{}).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (u userRepository) FindOneById(ctx context.Context, id *uint) (*domain.User, error) {
	var result *domain.User

	if err := u.gorm.Model(&domain.User{}).Where("id = ?", *id).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u userRepository) FindOneByEmail(ctx context.Context, email *string) (*domain.User, error) {
	var result *domain.User

	if err := u.gorm.Model(&domain.User{}).Where("email = ?", *email).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u userRepository) FindOneByUsername(ctx context.Context, username *string) (*domain.User, error) {
	var result *domain.User

	if err := u.gorm.Model(&domain.User{}).Where("username = ?", *username).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u userRepository) InsertOne(ctx context.Context, user *domain.User) (*domain.User, error) {
	if err := u.gorm.Model(&domain.User{}).Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (u userRepository) UpdateOneById(ctx context.Context, id *uint, user *domain.User) (*domain.User, error) {
	var userModel domain.User

	if err := u.gorm.Model(&userModel).Clauses(clause.Returning{}).Where("id = ?", *id).Updates(user).Error; err != nil {
		return nil, err
	}

	return &userModel, nil
}

func (u userRepository) DeleteOneById(ctx context.Context, id *uint) (*domain.User, error) {
	var userModel domain.User

	if err := u.gorm.Model(&domain.User{}).Clauses(clause.Returning{}).Where("id = ?", id).Delete(&userModel).Error; err != nil {
		return nil, err
	}

	return &userModel, nil
}
