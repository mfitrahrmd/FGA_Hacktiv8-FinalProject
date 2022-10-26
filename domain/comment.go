package domain

import "context"

type Comment struct {
	Base
	Message string `json:"message,omitempty"`
	UserId  uint   `json:"user_id,omitempty"`
	PhotoId uint   `json:"photo_id,omitempty"`
}

type CommentAdd struct {
	Message string `json:"message" binding:"required"`
	PhotoId uint   `json:"photo_id" binding:"required"`
}

type CommentUpdateData struct {
	Message string `json:"message" binding:"required"`
}

type CommentWithUserAndPhotoData struct {
	Comment
	User  *User  `json:"user,omitempty" gorm:"embedded"`
	Photo *Photo `json:"photo,omitempty" gorm:"embedded"`
}

type CommentRepository interface {
	FindAllWithUserAndPhotoData(ctx context.Context) (*[]CommentWithUserAndPhotoData, error)
	FindOneById(ctx context.Context, id *uint) (*Comment, error)
	FindALlByUserId(ctx context.Context, userId *uint) (*[]Comment, error)
	InsertOne(ctx context.Context, comment *Comment) (*Comment, error)
	UpdateOneById(ctx context.Context, id *uint, comment *Comment) (*Comment, error)
	DeleteOneById(ctx context.Context, id *uint) (*Comment, error)
}
