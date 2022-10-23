package socialMedia

import (
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	socialMediaUsecase "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/socialMedia"
	"net/http"
	"strconv"
)

type SocialMediaController interface {
	PostSocialMedia(ctx *gin.Context)
	GetAllSocialMedias(ctx *gin.Context)
	PutSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaController struct {
	socialMediaUsecase socialMediaUsecase.SocialMediaUsecase
}

func NewSocialMediaController(socialMediaUsecase socialMediaUsecase.SocialMediaUsecase) SocialMediaController {
	return socialMediaController{
		socialMediaUsecase: socialMediaUsecase,
	}
}

func (s socialMediaController) PostSocialMedia(ctx *gin.Context) {
	ctx.Set("userId", 1)
	getUserId := ctx.MustGet("userId").(int)
	userId := uint(getUserId)

	var bindSocialMedia domain.SocialMedia

	err := ctx.ShouldBindJSON(&bindSocialMedia)
	if err != nil {
		ctx.Abort()

		return
	}

	addedSocialMedia, err := s.socialMediaUsecase.AddSocialMedia(ctx, &userId, &bindSocialMedia)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusCreated, addedSocialMedia)
}

func (s socialMediaController) GetAllSocialMedias(ctx *gin.Context) {
	socialMedias, err := s.socialMediaUsecase.GetAllSocialMedias(ctx)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, socialMedias)
}

func (s socialMediaController) PutSocialMedia(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")
	conv, _ := strconv.ParseUint(socialMediaId, 10, 64)
	convSocialMediaId := uint(conv)

	ctx.Set("userId", 1)
	getUserId := ctx.MustGet("userId").(int)
	userId := uint(getUserId)

	var bindSocialMedia domain.SocialMedia

	err := ctx.ShouldBindJSON(&bindSocialMedia)
	if err != nil {
		ctx.Abort()

		return
	}

	updatedSocialMedia, err := s.socialMediaUsecase.UpdateSocialMedia(ctx, &userId, &convSocialMediaId, &bindSocialMedia)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, updatedSocialMedia)
}

func (s socialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId := ctx.Param("socialMediaId")
	conv, _ := strconv.ParseUint(socialMediaId, 10, 64)
	convSocialMediaId := uint(conv)

	ctx.Set("userId", 1)
	getUserId := ctx.MustGet("userId").(int)
	userId := uint(getUserId)

	_, err := s.socialMediaUsecase.DeleteSocialMedia(ctx, &userId, &convSocialMediaId)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
