package domain

import "context"

type SocialMedia struct {
	Base
	Name           string `json:"name,omitempty"`
	SocialMediaUrl string `json:"social_media_url,omitempty"`
	UserId         uint   `json:"user_id,omitempty"`
}

type SocialMediaWithUserData struct {
	SocialMedia
	User *User `json:"user,omitempty" gorm:"embedded"`
}

type SocialMediaRepository interface {
	FindAllWithUserData(ctx context.Context) (*[]SocialMediaWithUserData, error)
	FindOneById(ctx context.Context, id *uint) (*SocialMedia, error)
	FindALlByUserId(ctx context.Context, userId *uint) (*[]SocialMedia, error)
	InsertOne(ctx context.Context, socialMedia *SocialMedia) (*SocialMedia, error)
	UpdateOneById(ctx context.Context, id *uint, socialMedia *SocialMedia) (*SocialMedia, error)
	DeleteOneById(ctx context.Context, id *uint) (*SocialMedia, error)
}
