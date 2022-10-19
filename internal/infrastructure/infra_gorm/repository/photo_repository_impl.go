package repository

import (
	"context"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"gorm.io/gorm"
)

type photoRepository struct {
	gorm *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{
		gorm: db,
	}
}

func (p photoRepository) FindAll(ctx context.Context) (*[]domain.Photo, error) {
	var result *[]domain.Photo

	if err := p.gorm.Debug().Find(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (p photoRepository) InsertOne(ctx context.Context, photo *domain.Photo) error {
	if err := p.gorm.Debug().Model(&domain.Photo{}).Save(photo).Error; err != nil {
		return err
	}

	return nil
}

func (p photoRepository) UpdateOneById(ctx context.Context, id *string, photo *domain.Photo) error {
	if err := p.gorm.Debug().Model(&domain.Photo{}).Where("id = ?", *id).Updates(photo).Error; err != nil {
		return err
	}

	return nil
}

func (p photoRepository) DeleteOneById(ctx context.Context, id *string) error {
	if err := p.gorm.Debug().Model(&domain.Photo{}).Where("id = ?", *id).Delete(&domain.Photo{}).Error; err != nil {
		return err
	}

	return nil
}
