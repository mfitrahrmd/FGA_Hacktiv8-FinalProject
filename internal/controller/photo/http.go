package photo

import (
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/photo"
	"net/http"
	"strconv"
)

type PhotoController interface {
	PostPhoto(ctx *gin.Context)
	GetAllPhotos(ctx *gin.Context)
	PutPhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoHttp struct {
	photoUsecase photo.PhotoUsecase
}

func NewUserHttp(photoUsecase photo.PhotoUsecase) PhotoController {
	return &photoHttp{
		photoUsecase: photoUsecase,
	}
}

func (p photoHttp) PostPhoto(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint)

	var bindPhoto domain.Photo

	err := ctx.ShouldBindJSON(&bindPhoto)
	if err != nil {
		ctx.Abort()

		return
	}

	addedPhoto, err := p.photoUsecase.AddPhoto(ctx, &userId, &bindPhoto)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusCreated, addedPhoto)
}

func (p photoHttp) GetAllPhotos(ctx *gin.Context) {
	photos, err := p.photoUsecase.GetAllPhotos(ctx)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (p photoHttp) PutPhoto(ctx *gin.Context) {
	paramPhotoId := ctx.Param("photoId")
	conv, _ := strconv.ParseUint(paramPhotoId, 10, 64)
	photoId := uint(conv)

	userId := ctx.MustGet("userId").(uint)

	var bindPhoto domain.Photo

	err := ctx.ShouldBindJSON(&bindPhoto)
	if err != nil {
		ctx.Abort()

		return
	}

	updatedPhoto, err := p.photoUsecase.UpdatePhoto(ctx, &userId, &photoId, &bindPhoto)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, updatedPhoto)
}

func (p photoHttp) DeletePhoto(ctx *gin.Context) {
	paramPhotoId := ctx.Param("photoId")
	conv, _ := strconv.ParseUint(paramPhotoId, 10, 64)
	photoId := uint(conv)

	userId := ctx.MustGet("userId").(uint)

	_, err := p.photoUsecase.DeletePhoto(ctx, &userId, &photoId)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	})
}
