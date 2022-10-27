package domain

import "context"

type SocialMedia struct {
	Base
	Name           string `json:"name,omitempty"`
	SocialMediaUrl string `json:"social_media_url,omitempty"`
	UserId         uint   `json:"user_id,omitempty"`
}

type SocialMediaAdd struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
}

type SocialMediaAddResponse struct {
	Base
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         uint   `json:"user_id"`
}

type SocialMediaUpdateData struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaUrl string `json:"social_media_url" binding:"required"`
}

type SocialMediaUpdateDataResponse struct {
	Base
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
	UserId         uint   `json:"user_id"`
}

type SocialMediaWithUserData struct {
	SocialMedia
	User *User `json:"user,omitempty" gorm:"embedded"`
}

type SocialMediaDeleteResponse struct {
	Message string `json:"message"`
}

type SocialMediaRepository interface {
	FindAllWithUserData(ctx context.Context) (*[]SocialMediaWithUserData, error)
	FindOneById(ctx context.Context, id *uint) (*SocialMedia, error)
	FindALlByUserId(ctx context.Context, userId *uint) (*[]SocialMedia, error)
	InsertOne(ctx context.Context, socialMedia *SocialMedia) (*SocialMedia, error)
	UpdateOneById(ctx context.Context, id *uint, socialMedia *SocialMedia) (*SocialMedia, error)
	DeleteOneById(ctx context.Context, id *uint) (*SocialMedia, error)
}
