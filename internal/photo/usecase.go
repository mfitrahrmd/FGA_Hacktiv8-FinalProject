package photo

import (
	"context"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
)

type photoUsecase struct {
	photoRespository domain.PhotoRepository
}

func (p photoUsecase) AddPhoto(ctx context.Context, photo *domain.Photo) error {
	err := p.photoRespository.InsertOne(ctx, photo)
	if err != nil {
		return err
	}

	return nil
}
