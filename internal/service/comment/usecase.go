package comment

import (
	"context"
	"errors"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service"
	"gorm.io/gorm"
)

type CommentUsecase interface {
	AddComment(ctx context.Context, userId *uint, comment *domain.Comment) (*domain.Comment, error)
	GetAllComments(ctx context.Context) (*[]domain.CommentWithUserAndPhotoData, error)
	UpdateComment(ctx context.Context, commentId *uint, userId *uint, comment *domain.Comment) (*domain.Comment, error)
	DeleteComment(ctx context.Context, commentId *uint, userId *uint) (*domain.Comment, error)
}

type commentUsecase struct {
	commentRepository domain.CommentRepository
	userRepository    domain.UserRepository
	photoRepository   domain.PhotoRepository
}

func NewCommentUsecase(commentRepository domain.CommentRepository, userRepository domain.UserRepository, photoRepository domain.PhotoRepository) CommentUsecase {
	return &commentUsecase{
		commentRepository: commentRepository,
		userRepository:    userRepository,
		photoRepository:   photoRepository,
	}
}

func (c commentUsecase) VerifyCommentOwner(ctx context.Context, commentId *uint, userId *uint) (bool, error) {
	foundComment, err := c.commentRepository.FindOneById(ctx, commentId)
	if err != nil {
		return false, err
	}

	if foundComment.PhotoId != *userId {
		return false, nil
	}

	return true, nil
}

func (c commentUsecase) AddComment(ctx context.Context, userId *uint, comment *domain.Comment) (*domain.Comment, error) {
	comment.UserId = *userId

	addedComment, err := c.commentRepository.InsertOne(ctx, comment)
	if err != nil {
		return nil, err
	}

	return addedComment, nil
}

func (c commentUsecase) GetAllComments(ctx context.Context) (*[]domain.CommentWithUserAndPhotoData, error) {
	foundComments, err := c.commentRepository.FindAllWithUserAndPhotoData(ctx)
	if err != nil {
		return nil, err
	}

	return foundComments, nil
}

func (c commentUsecase) UpdateComment(ctx context.Context, commentId *uint, userId *uint, comment *domain.Comment) (*domain.Comment, error) {
	foundComment, err := c.commentRepository.FindOneById(ctx, commentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.NewServiceError(service.NOT_FOUND)
		}

		return nil, err
	}

	isVerified, err := c.VerifyCommentOwner(ctx, &foundComment.Id, userId)
	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, service.NewServiceError(service.ACCESS_DENIED)
	}

	updatedComment, err := c.commentRepository.UpdateOneById(ctx, commentId, comment)
	if err != nil {
		return nil, err
	}

	return updatedComment, nil
}

func (c commentUsecase) DeleteComment(ctx context.Context, commentId *uint, userId *uint) (*domain.Comment, error) {
	foundComment, err := c.commentRepository.FindOneById(ctx, commentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service.NewServiceError(service.NOT_FOUND)
		}

		return nil, err
	}

	isVerified, err := c.VerifyCommentOwner(ctx, &foundComment.Id, userId)
	if err != nil {
		return nil, err
	}

	if !isVerified {
		return nil, service.NewServiceError(service.ACCESS_DENIED)
	}

	deletedComment, err := c.commentRepository.DeleteOneById(ctx, commentId)
	if err != nil {
		return nil, err
	}

	return deletedComment, nil
}
