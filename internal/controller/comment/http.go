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

func (c commentHttp) GetAllComments(ctx *gin.Context) {
	comments, err := c.commentUsecase.GetAllComments(c.ctx)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, comments)
}

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
