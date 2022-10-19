package domain

import (
	"context"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/infrastructure/infra_uuid"
	"gorm.io/gorm"
)

type Photo struct {
	Base
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl"`
	UserId   string `json:"userId"`
}

type PhotoRepository interface {
	FindAll(ctx context.Context) (*[]Photo, error)
	InsertOne(ctx context.Context, photo *Photo) error
	UpdateOneById(ctx context.Context, id *string, photo *Photo) error
	DeleteOneById(ctx context.Context, id *string) error
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	// generate user id
	p.Id = infra_uuid.GenerateUUID("photo")

	return nil
}
