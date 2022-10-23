package photo

import (
	"context"
	"errors"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service"
	"gorm.io/gorm"
)

type PhotoUsecase interface {
	AddPhoto(ctx context.Context, userId *uint, photo *domain.Photo) (*domain.Photo, error)
	GetAllPhotos(ctx context.Context) (*[]domain.PhotoWithUserData, error)
	UpdatePhoto(ctx context.Context, userId *uint, photoId *uint, photo *domain.Photo) (*domain.Photo, error)
	DeletePhoto(ctx context.Context, userId *uint, photoId *uint) (*domain.Photo, error)
}

type photoUsecase struct {
	photoRepository domain.PhotoRepository
}

func NewPhotoUsecase(photoRepository domain.PhotoRepository) PhotoUsecase {
	return &photoUsecase{
		photoRepository: photoRepository,
	}
}

func (p photoUsecase) VerifyPhotoOwner(ctx context.Context, photoId *uint, userId *uint) (bool, error) {
	foundPhoto, err := p.photoRepository.FindOneById(ctx, photoId)
	if err != nil {
		return false, err
	}

	if foundPhoto.UserId != *userId {
		return false, nil
	}

	return true, nil
}

func (p photoUsecase) AddPhoto(ctx context.Context, userId *uint, photo *domain.Photo) (*domain.Photo, error) {
	photo.UserId = *userId

	addedPhoto, err := p.photoRepository.InsertOne(ctx, photo)
	if err != nil {
		return nil, err
	}

	return addedPhoto, nil
}

func (p photoUsecase) GetAllPhotos(ctx context.Context) (*[]domain.PhotoWithUserData, error) {
	foundPhotos, err := p.photoRepository.FindAllWithUserData(ctx)
	if err != nil {
		return nil, err
	}

	return foundPhotos, nil
}

func (p photoUsecase) UpdatePhoto(ctx context.Context, userId *uint, photoId *uint, photo *domain.Photo) (*domain.Photo, error) {
	_, err := p.photoRepository.FindOneById(ctx, photoId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.NewServiceError(service.NOT_FOUND)
		}

		return nil, err
	}

	isVerified, err := p.VerifyPhotoOwner(ctx, photoId, userId)
	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, service.NewServiceError(service.ACCESS_DENIED)
	}

	updatedPhoto, err := p.photoRepository.UpdateOneById(ctx, photoId, photo)
	if err != nil {
		return nil, err
	}

	return updatedPhoto, nil
}

func (p photoUsecase) DeletePhoto(ctx context.Context, userId *uint, photoId *uint) (*domain.Photo, error) {
	foundPhoto, err := p.photoRepository.FindOneById(ctx, photoId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.NewServiceError(service.NOT_FOUND)
		}

		return nil, err
	}

	isVerified, err := p.VerifyPhotoOwner(ctx, &foundPhoto.Id, userId)
	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, service.NewServiceError(service.ACCESS_DENIED)
	}

	deletedPhoto, err := p.photoRepository.DeleteOneById(ctx, photoId)
	if err != nil {
		return nil, err
	}

	return deletedPhoto, nil
}
