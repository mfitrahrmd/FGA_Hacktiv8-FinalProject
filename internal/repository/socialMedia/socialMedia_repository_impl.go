package socialMedia

import (
	"context"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type socialMediaRepository struct {
	gorm *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{
		gorm: db,
	}
}

func (s socialMediaRepository) FindAllWithUserData(ctx context.Context) (*[]domain.SocialMediaWithUserData, error) {
	var result []domain.SocialMediaWithUserData

	err := s.gorm.Model(&domain.SocialMedia{}).Select("social_media.*, users.id, users.username").Joins("JOIN users ON social_media.user_id = users.id").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s socialMediaRepository) FindOneById(ctx context.Context, id *uint) (*domain.SocialMedia, error) {
	var result domain.SocialMedia

	if err := s.gorm.Model(&domain.SocialMedia{}).Where("id = ?", *id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (s socialMediaRepository) FindALlByUserId(ctx context.Context, userId *uint) (*[]domain.SocialMedia, error) {
	//TODO implement me
	panic("implement me")
}

func (s socialMediaRepository) InsertOne(ctx context.Context, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	if err := s.gorm.Model(&domain.SocialMedia{}).Save(socialMedia).Error; err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (s socialMediaRepository) UpdateOneById(ctx context.Context, id *uint, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	var socialMediaModel domain.SocialMedia

	if err := s.gorm.Model(&socialMediaModel).Clauses(clause.Returning{}).Where("id = ?", *id).Updates(socialMedia).Error; err != nil {
		return nil, err
	}

	return &socialMediaModel, nil
}

func (s socialMediaRepository) DeleteOneById(ctx context.Context, id *uint) (*domain.SocialMedia, error) {
	var socialMediaModel domain.SocialMedia

	if err := s.gorm.Model(&domain.SocialMedia{}).Clauses(clause.Returning{}).Where("id = ?", *id).Delete(&socialMediaModel).Error; err != nil {
		return nil, err
	}

	return &socialMediaModel, nil
}
