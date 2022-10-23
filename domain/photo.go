package domain

import (
	"context"
)

type Photo struct {
	Base
	Title    string    `json:"title,omitempty"`
	Caption  string    `json:"caption,omitempty"`
	PhotoUrl string    `json:"photo_url,omitempty"`
	UserId   uint      `json:"user_id,omitempty"`
	Comments []Comment `json:"comment,omitempty"`
}

type PhotoWithUserData struct {
	Photo
	User *User `json:"user,omitempty" gorm:"embedded"`
}

type PhotoRepository interface {
	FindAllWithUserData(ctx context.Context) (*[]PhotoWithUserData, error)
	FindOneById(ctx context.Context, id *uint) (*Photo, error)
	FindALlByUserId(ctx context.Context, userId *uint) (*[]Photo, error)
	InsertOne(ctx context.Context, photo *Photo) (*Photo, error)
	UpdateOneById(ctx context.Context, id *uint, photo *Photo) (*Photo, error)
	DeleteOneById(ctx context.Context, id *uint) (*Photo, error)
}
