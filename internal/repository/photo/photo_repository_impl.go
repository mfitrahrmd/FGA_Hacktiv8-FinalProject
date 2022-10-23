package photo

import (
	"context"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type photoRepository struct {
	gorm *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{
		gorm: db,
	}
}

func (p photoRepository) FindAllWithUserData(ctx context.Context) (*[]domain.PhotoWithUserData, error) {
	var result []domain.PhotoWithUserData

	if err := p.gorm.Model(&domain.Photo{}).Select("photos.*, users.email, users.username").Joins("JOIN users ON photos.user_id = users.id").Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (p photoRepository) FindOneById(ctx context.Context, id *uint) (*domain.Photo, error) {
	var result domain.Photo

	if err := p.gorm.Model(&domain.Photo{}).Where("id = ?", *id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (p photoRepository) FindALlByUserId(ctx context.Context, userId *uint) (*[]domain.Photo, error) {
	var result []domain.Photo

	if err := p.gorm.Model(&domain.Photo{}).Where("user_id = ?", *userId).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (p photoRepository) InsertOne(ctx context.Context, photo *domain.Photo) (*domain.Photo, error) {
	if err := p.gorm.Model(&domain.Photo{}).Save(photo).Error; err != nil {
		return nil, err
	}

	return photo, nil
}

func (p photoRepository) UpdateOneById(ctx context.Context, id *uint, photo *domain.Photo) (*domain.Photo, error) {
	var photoModel domain.Photo

	if err := p.gorm.Model(&photoModel).Clauses(clause.Returning{}).Where("id = ?", *id).Updates(photo).Error; err != nil {
		return nil, err
	}

	return &photoModel, nil
}

func (p photoRepository) DeleteOneById(ctx context.Context, id *uint) (*domain.Photo, error) {
	var photoModel domain.Photo

	if err := p.gorm.Model(&domain.Photo{}).Clauses(clause.Returning{}).Where("id = ?", *id).Delete(&photoModel).Error; err != nil {
		return nil, err
	}

	return &photoModel, nil
}
