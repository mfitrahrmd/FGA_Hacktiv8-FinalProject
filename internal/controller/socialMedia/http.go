package socialMedia

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	socialMediaUsecase "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/socialMedia"
)

type SocialMediaController interface {
	PostSocialMedia(ctx *gin.Context)
	GetAllSocialMedias(ctx *gin.Context)
	PutSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaController struct {
	socialMediaUsecase socialMediaUsecase.SocialMediaUsecase
	ctx                context.Context
}

func NewSocialMediaController(ctx context.Context, socialMediaUsecase socialMediaUsecase.SocialMediaUsecase) SocialMediaController {
	return socialMediaController{
		socialMediaUsecase: socialMediaUsecase,
		ctx:                ctx,
	}
}

// @Summary Add user's social media
// @Tags User's Social Media
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Param photo body domain.SocialMediaAdd true "Social Media Data"
// @Success 201 {object} domain.SocialMediaAddResponse
// @Failure 400 {object} middleware.Error
// @Failure 401 {object} middleware.Error
// @Failure 500 {object} middleware.Error
// @Router /socialmedias [post]
func (s socialMediaController) PostSocialMedia(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint)

	var bindSocialMedia domain.SocialMediaAdd

	err := ctx.ShouldBindJSON(&bindSocialMedia)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	addedSocialMedia, err := s.socialMediaUsecase.AddSocialMedia(s.ctx, &userId, &domain.SocialMedia{
		Name:           bindSocialMedia.Name,
		SocialMediaUrl: bindSocialMedia.SocialMediaUrl,
	})
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusCreated, addedSocialMedia)
}

// @Summary Get all social medias
// @Tags User's Social Media
// @Produce json
// @Param Authorization header string true "Access Token"
// @Success 200 {object} domain.SocialMediaWithUserData
// @Failure 401 {object} middleware.Error
// @Failure 500 {object} middleware.Error
// @Router /socialmedias [get]
func (s socialMediaController) GetAllSocialMedias(ctx *gin.Context) {
	socialMedias, err := s.socialMediaUsecase.GetAllSocialMedias(s.ctx)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, socialMedias)
}

// @Summary Update existing user's social media
// @Tags User's Social Media
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Param socialMediaId path int true "Social Media ID"
// @Param photo body domain.SocialMediaUpdateData true "Social Media Data"
// @Success 200 {object} domain.SocialMediaUpdateDataResponse
// @Failure 400 {object} middleware.Error
// @Failure 401 {object} middleware.Error
// @Failure 404 {object} middleware.Error
// @Failure 500 {object} middleware.Error
// @Router /socialmedias/{socialMediaId} [put]
func (s socialMediaController) PutSocialMedia(ctx *gin.Context) {
	paramSocialMediaId := ctx.Param("socialMediaId")
	conv, _ := strconv.ParseUint(paramSocialMediaId, 10, 64)
	socialMediaId := uint(conv)

	userId := ctx.MustGet("userId").(uint)

	var bindSocialMedia domain.SocialMediaUpdateData

	err := ctx.ShouldBindJSON(&bindSocialMedia)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	updatedSocialMedia, err := s.socialMediaUsecase.UpdateSocialMedia(s.ctx, &userId, &socialMediaId, &domain.SocialMedia{
		Name:           bindSocialMedia.Name,
		SocialMediaUrl: bindSocialMedia.SocialMediaUrl,
	})
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, updatedSocialMedia)
}

// @Summary Delete existing user's social media
// @Tags User's Social Media
// @Produce json
// @Param Authorization header string true "Access Token"
// @Param socialMediaId path int true "Social Media ID"
// @Success 200 {object} domain.SocialMediaUpdateDataResponse
// @Failure 401 {object} middleware.Error
// @Failure 404 {object} middleware.Error
// @Failure 500 {object} middleware.Error
// @Router /socialmedias/{socialMediaId} [delete]
func (s socialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	paramSocialMediaId := ctx.Param("socialMediaId")
	conv, _ := strconv.ParseUint(paramSocialMediaId, 10, 64)
	socialMediaId := uint(conv)

	userId := ctx.MustGet("userId").(uint)

	_, err := s.socialMediaUsecase.DeleteSocialMedia(s.ctx, &userId, &socialMediaId)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
