package comment

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/comment"
)

type CommentController interface {
	PostComment(ctx *gin.Context)
	GetAllComments(ctx *gin.Context)
	PutComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentHttp struct {
	commentUsecase comment.CommentUsecase
	ctx            context.Context
}

func NewCommentHttp(ctx context.Context, commentUsecase comment.CommentUsecase) CommentController {
	return &commentHttp{
		commentUsecase: commentUsecase,
		ctx:            ctx,
	}
}

// @Summary Add comment to existing photo
// @Tags Photo's Comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Param comment body domain.CommentAdd true "Comment Data"
// @Success 201 {object} domain.CommentAddResponse
// @Failure 400 {object} middleware.Error
// @Failure 401 {object} middleware.Error
// @Failure 500 {object} middleware.Error
// @Router /comments [post]
func (c commentHttp) PostComment(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint)

	var bindComment domain.CommentAdd

	err := ctx.ShouldBindJSON(&bindComment)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	addedComment, err := c.commentUsecase.AddComment(c.ctx, &userId, &domain.Comment{
		Message: bindComment.Message,
		PhotoId: bindComment.PhotoId,
	})
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusCreated, addedComment)
}

// @Summary Get all comments
// @Tags Photo's Comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Success 200 {object} domain.CommentWithUserAndPhotoData
// @Failure 401 {object} middleware.Error
// @Failure 500 {object} middleware.Error
// @Router /comments [get]
func (c commentHttp) GetAllComments(ctx *gin.Context) {
	comments, err := c.commentUsecase.GetAllComments(c.ctx)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// @Summary Update existing comment
// @Tags Photo's Comments
// @Accept json
// @Produce json
// @Param Authorization header string true "Access Token"
// @Param commentId path int true "Comment ID"
// @Param comment body domain.CommentUpdateData true "Comment Data"
// @Success 200 {object} domain.CommentUpdateDataResponse
// @Failure 400 {object} middleware.Error
// @Failure 401 {object} middleware.Error
// @Failure 404 {object} middleware.Error
// @Failure 500 {object} middleware.Error
// @Router /comments/{commentId} [put]
func (c commentHttp) PutComment(ctx *gin.Context) {
	paramCommentId := ctx.Param("commentId")
	conv, _ := strconv.ParseUint(paramCommentId, 10, 64)
	commentId := uint(conv)

	userId := ctx.MustGet("userId").(uint)

	var bindComment domain.CommentUpdateData

	err := ctx.ShouldBindJSON(&bindComment)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	updatedComment, err := c.commentUsecase.UpdateComment(c.ctx, &commentId, &userId, &domain.Comment{
		Message: bindComment.Message,
	})
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, updatedComment)
}

// @Summary Delete existing comment
// @Tags Photo's Comments
// @Produce json
// @Param Authorization header string true "Access Token"
// @Param commentId path int true "Comment ID"
// @Success 200 {object} domain.CommentDeleteResponse
// @Failure 401 {object} middleware.Error
// @Failure 404 {object} middleware.Error
// @Failure 500 {object} middleware.Error
// @Router /comments/{commentId} [delete]
func (c commentHttp) DeleteComment(ctx *gin.Context) {
	paramPhotoId := ctx.Param("commentId")
	conv, _ := strconv.ParseUint(paramPhotoId, 10, 64)
	commentId := uint(conv)

	userId := ctx.MustGet("userId").(uint)

	_, err := c.commentUsecase.DeleteComment(c.ctx, &commentId, &userId)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
