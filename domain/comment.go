package domain

import "context"

type Comment struct {
	Base
	Message string `json:"message,omitempty"`
	UserId  uint   `json:"user_id,omitempty"`
	PhotoId uint   `json:"photo_id,omitempty"`
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
