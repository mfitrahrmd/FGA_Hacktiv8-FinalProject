package socialMedia

import (
	"context"
	"errors"
	"fmt"

	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service"
	"gorm.io/gorm"
)

type SocialMediaUsecase interface {
	AddSocialMedia(ctx context.Context, userId *uint, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error)
	GetAllSocialMedias(ctx context.Context) (*[]domain.SocialMediaWithUserData, error)
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

	if foundPhoto.UserId != *userId {
		return false, nil
	}

	return true, nil
}

func (s socialMediaUsecase) AddSocialMedia(ctx context.Context, userId *uint, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	socialMedia.UserId = *userId
	fmt.Println(socialMedia)

	addedSocialMedia, err := s.socialMediaRepository.InsertOne(ctx, socialMedia)
	if err != nil {
		return nil, err
	}

	return addedSocialMedia, err
}

func (s socialMediaUsecase) GetAllSocialMedias(ctx context.Context) (*[]domain.SocialMediaWithUserData, error) {
	foundSocialMedias, err := s.socialMediaRepository.FindAllWithUserData(ctx)
	if err != nil {
		return nil, err
	}

	return foundSocialMedias, nil
}

func (s socialMediaUsecase) UpdateSocialMedia(ctx context.Context, userId *uint, socialMediaId *uint, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	_, err := s.socialMediaRepository.FindOneById(ctx, socialMediaId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.SOCIAL_MEDIA_NOT_FOUND
		}

		return nil, err
	}

	isVerified, err := s.VerifySocialMediaOwner(ctx, socialMediaId, userId)
	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, service.ACCESS_DENIED
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
			return nil, service.SOCIAL_MEDIA_NOT_FOUND
		}

		return nil, err
	}

	isVerified, err := s.VerifySocialMediaOwner(ctx, socialMediaId, userId)
	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, service.ACCESS_DENIED
	}

	deletedSocialMedia, err := s.socialMediaRepository.DeleteOneById(ctx, socialMediaId)
	if err != nil {
		return nil, err
	}

	return deletedSocialMedia, nil
}
