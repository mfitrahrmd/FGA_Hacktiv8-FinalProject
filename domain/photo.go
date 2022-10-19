package domain

import (
	"context"
)

type Photo struct {
	Base
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl"`
	UserId   uint   `json:"userId"`
}

type PhotoRepository interface {
	FindOneById(ctx context.Context, id *uint) (*Photo, error)
	FindALlByUserId(ctx context.Context, userId *uint) (*[]Photo, error)
	InsertOne(ctx context.Context, photo *Photo) (*Photo, error)
	UpdateOneById(ctx context.Context, id *uint, photo *Photo) (*Photo, error)
	DeleteOneById(ctx context.Context, id *uint) (*Photo, error)
}
