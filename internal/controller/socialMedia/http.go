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

func (s socialMediaController) GetAllSocialMedias(ctx *gin.Context) {
	socialMedias, err := s.socialMediaUsecase.GetAllSocialMedias(s.ctx)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, socialMedias)
}

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
