package photo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/user"
	"gorm.io/gorm"
)

type photoUsecase struct {
	photoRepository domain.PhotoRepository
}

func NewPhotoUsecase(photoRepository domain.PhotoRepository) *photoUsecase {
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

func (p photoUsecase) AddUserPhoto(ctx context.Context, photo *domain.Photo) (*domain.Photo, error) {
	userTokenPayload := ctx.Value("userTokenPayload").(user.TokenPayload)

	photo.UserId = userTokenPayload.Id

	addedPhoto, err := p.photoRepository.InsertOne(ctx, photo)
	if err != nil {
		return nil, err
	}

	return addedPhoto, nil
}

func (p photoUsecase) GetAllUserPhotos(ctx context.Context) (*[]map[string]any, error) {
	userTokenPayload := ctx.Value("userTokenPayload").(user.TokenPayload)

	foundPhotos, err := p.photoRepository.FindALlByUserId(ctx, &userTokenPayload.Id)
	if err != nil {
		return nil, err
	}

	var allPhotos []map[string]any

	if len(*foundPhotos) > 0 {
		data, _ := json.Marshal(foundPhotos)
		json.Unmarshal(data, &allPhotos)

		for _, photo := range allPhotos {
			photo["User"] = map[string]any{
				"email":    userTokenPayload.Email,
				"username": userTokenPayload.Username,
			}
		}
	}

	return &allPhotos, nil
}

func (p photoUsecase) UpdateUserPhoto(ctx context.Context, id *uint, photo *domain.Photo) (*domain.Photo, error) {
	userTokenPayload := ctx.Value("userTokenPayload").(user.TokenPayload)

	foundPhoto, err := p.photoRepository.FindOneById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewPhotoError(PHOTO_NOT_FOUND)
		}

		return nil, err
	}

	isVerified, err := p.VerifyPhotoOwner(ctx, &foundPhoto.UserId, &userTokenPayload.Id)
	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, NewPhotoError(ACCESS_DENIED)
	}

	updatedPhoto, err := p.photoRepository.UpdateOneById(ctx, id, photo)
	if err != nil {
		return nil, err
	}

	return updatedPhoto, nil
}

func (p photoUsecase) DeleteUserPhoto(ctx context.Context, id *uint) (*domain.Photo, error) {
	userTokenPayload := ctx.Value("userTokenPayload").(user.TokenPayload)

	foundPhoto, err := p.photoRepository.FindOneById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, NewPhotoError(PHOTO_NOT_FOUND)
		}

		return nil, err
	}

	isVerified, err := p.VerifyPhotoOwner(ctx, &foundPhoto.Id, &userTokenPayload.Id)
	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, NewPhotoError(ACCESS_DENIED)
	}

	deletedPhoto, err := p.photoRepository.DeleteOneById(ctx, id)
	if err != nil {
		return nil, err
	}

	return deletedPhoto, nil
}
