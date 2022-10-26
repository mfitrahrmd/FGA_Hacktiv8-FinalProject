package comment

import (
	"context"

	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type commentRepository struct {
	gorm *gorm.DB
}

func NewCommentRepository(db *gorm.DB) domain.CommentRepository {
	return &commentRepository{
		gorm: db,
	}
}

func (c commentRepository) FindAllWithUserAndPhotoData(ctx context.Context) (*[]domain.CommentWithUserAndPhotoData, error) {
	var result []domain.CommentWithUserAndPhotoData

	rows, err := c.gorm.Model(&domain.Comment{}).Select("comments.id, comments.message, comments.photo_id, comments.user_id, comments.created_at, comments.updated_at, users.id, users.email, users.username, photos.id, photos.title, photos.caption, photos.photo_url, photos.user_id").Joins("JOIN users ON comments.user_id = users.id JOIN photos ON comments.photo_id = photos.id").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r domain.CommentWithUserAndPhotoData
		var user domain.User
		var photo domain.Photo
		rows.Scan(&r.Id, &r.Message, &r.PhotoId, &r.UserId, &r.CreatedAt, &r.UpdatedAt, &user.Id, &user.Email, &user.Username, &photo.Id, &photo.Title, &photo.Caption, &photo.PhotoUrl, &photo.UserId)
		r.User = &user
		r.Photo = &photo
		result = append(result, r)
	}

	return &result, nil
}

func (c commentRepository) FindOneById(ctx context.Context, id *uint) (*domain.Comment, error) {
	var result domain.Comment

	if err := c.gorm.Model(&domain.Comment{}).Where("id = ?", *id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (c commentRepository) FindALlByUserId(ctx context.Context, userId *uint) (*[]domain.Comment, error) {
	var result []domain.Comment

	if err := c.gorm.Model(&domain.Comment{}).Where("user_id = ?", *userId).Find(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (c commentRepository) InsertOne(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	if err := c.gorm.Model(&domain.Comment{}).Save(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (c commentRepository) UpdateOneById(ctx context.Context, id *uint, comment *domain.Comment) (*domain.Comment, error) {
	var commentModel domain.Comment

	if err := c.gorm.Model(&commentModel).Clauses(clause.Returning{}).Where("id = ?", *id).Updates(comment).Error; err != nil {
		return nil, err
	}

	return &commentModel, nil
}

func (c commentRepository) DeleteOneById(ctx context.Context, id *uint) (*domain.Comment, error) {
	var commentModel domain.Comment

	if err := c.gorm.Model(&domain.Comment{}).Clauses(clause.Returning{}).Where("id = ?", *id).Delete(&commentModel).Error; err != nil {
		return nil, err
	}

	return &commentModel, nil
}
