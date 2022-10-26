package socialMedia

import (
	"context"

	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type socialMediaRepository struct {
	gorm *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) domain.SocialMediaRepository {
	return &socialMediaRepository{
		gorm: db,
	}
}

func (s socialMediaRepository) FindAllWithUserData(ctx context.Context) (*[]domain.SocialMediaWithUserData, error) {
	var result []domain.SocialMediaWithUserData

	rows, err := s.gorm.Debug().Raw("SELECT social_media.id, social_media.name, social_media.social_media_url, social_media.user_id, social_media.created_at, social_media.updated_at, users.id, users.username, photos.photo_url FROM social_media JOIN users ON social_media.user_id = users.id LEFT JOIN photos ON users.id = photos.user_id").Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r domain.SocialMediaWithUserData
		var u domain.User
		var p domain.Photo
		rows.Scan(&r.Id, &r.Name, &r.SocialMediaUrl, &r.UserId, &r.CreatedAt, &r.UpdatedAt, &u.Id, &u.Username, &p.PhotoUrl)
		r.User = &u
		r.User.Photo = &p
		result = append(result, r)
	}

	return &result, nil
}

func (s socialMediaRepository) FindOneById(ctx context.Context, id *uint) (*domain.SocialMedia, error) {
	var result domain.SocialMedia

	if err := s.gorm.Model(&domain.SocialMedia{}).Where("id = ?", *id).First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

func (s socialMediaRepository) FindALlByUserId(ctx context.Context, userId *uint) (*[]domain.SocialMedia, error) {
	//TODO implement me
	panic("implement me")
}

func (s socialMediaRepository) InsertOne(ctx context.Context, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	if err := s.gorm.Model(&domain.SocialMedia{}).Save(socialMedia).Error; err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (s socialMediaRepository) UpdateOneById(ctx context.Context, id *uint, socialMedia *domain.SocialMedia) (*domain.SocialMedia, error) {
	var socialMediaModel domain.SocialMedia

	if err := s.gorm.Model(&socialMediaModel).Clauses(clause.Returning{}).Where("id = ?", *id).Updates(socialMedia).Error; err != nil {
		return nil, err
	}

	return &socialMediaModel, nil
}

func (s socialMediaRepository) DeleteOneById(ctx context.Context, id *uint) (*domain.SocialMedia, error) {
	var socialMediaModel domain.SocialMedia

	if err := s.gorm.Model(&domain.SocialMedia{}).Clauses(clause.Returning{}).Where("id = ?", *id).Delete(&socialMediaModel).Error; err != nil {
		return nil, err
	}

	return &socialMediaModel, nil
}
