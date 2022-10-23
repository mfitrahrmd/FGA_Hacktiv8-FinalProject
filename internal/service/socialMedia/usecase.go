package socialMedia

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service"
	"gorm.io/gorm"
)

type SocialMediaUsecase interface {
	AddSocialMedia(ctx context.Context, userId *uint, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error)
	GetAllSocialMedias(ctx context.Context) ([]map[string]any, error)
	UpdateSocialMedia(ctx context.Context, userId *uint, socialMediaId *uint, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, userId *uint, socialMediaId *uint) (*domain.SocialMedia, error)
}

type socialMediaUsecase struct {
	socialMediaRepository domain.SocialMediaRepository
	photoRepository       domain.PhotoRepository
}

func NewSocialMediaUsecase(socialMediaRepository domain.SocialMediaRepository, photoRepository domain.PhotoRepository) SocialMediaUsecase {
	return &socialMediaUsecase{
		socialMediaRepository: socialMediaRepository,
		photoRepository:       photoRepository,
	}
}

func (s socialMediaUsecase) VerifySocialMediaOwner(ctx context.Context, socialMediaId *uint, userId *uint) (bool, error) {
	foundPhoto, err := s.socialMediaRepository.FindOneById(ctx, socialMediaId)
	if err != nil {
		return false, err
	}

	return foundPhoto.UserId == *userId, nil
}

func (s socialMediaUsecase) AddSocialMedia(ctx context.Context, userId *uint, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	socialMedia.UserId = *userId

	addedSocialMedia, err := s.socialMediaRepository.InsertOne(ctx, socialMedia)
	if err != nil {
		return nil, err
	}

	return addedSocialMedia, err
}

func (s socialMediaUsecase) GetAllSocialMedias(ctx context.Context) ([]map[string]any, error) {
	foundSocialMedias, err := s.socialMediaRepository.FindAllWithUserData(ctx)
	if err != nil {
		return nil, err
	}

	socialMediasJson, _ := json.Marshal(foundSocialMedias)

	var result []map[string]any

	json.Unmarshal(socialMediasJson, &result)

	for _, r := range result {
		photoId := r["user"].(map[string]any)["id"].(float64)
		uintPhotoId := uint(photoId)

		foundUserPhoto, err := s.photoRepository.FindOneById(ctx, &uintPhotoId)
		if err != nil {
			return nil, err
		}

		r["user"].(map[string]any)["profile_image_url"] = foundUserPhoto.PhotoUrl
	}

	return result, nil
}

func (s socialMediaUsecase) UpdateSocialMedia(ctx context.Context, userId *uint, socialMediaId *uint, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	_, err := s.socialMediaRepository.FindOneById(ctx, socialMediaId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.NewServiceError(service.NOT_FOUND)
		}

		return nil, err
	}

	isVerified, err := s.VerifySocialMediaOwner(ctx, socialMediaId, userId)
	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, service.NewServiceError(service.ACCESS_DENIED)
	}

	updatedSocialMedia, err := s.socialMediaRepository.UpdateOneById(ctx, socialMediaId, socialMedia)
	if err != nil {
		return nil, err
	}

	return updatedSocialMedia, nil
}

func (s socialMediaUsecase) DeleteSocialMedia(ctx context.Context, userId *uint, socialMediaId *uint) (*domain.SocialMedia, error) {
	_, err := s.socialMediaRepository.FindOneById(ctx, socialMediaId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.NewServiceError(service.NOT_FOUND)
		}

		return nil, err
	}

	isVerified, err := s.VerifySocialMediaOwner(ctx, socialMediaId, userId)
	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, service.NewServiceError(service.ACCESS_DENIED)
	}

	deletedSocialMedia, err := s.socialMediaRepository.DeleteOneById(ctx, socialMediaId)
	if err != nil {
		return nil, err
	}

	return deletedSocialMedia, nil
}
