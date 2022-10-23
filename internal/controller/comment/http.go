package comment

import (
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/comment"
	"net/http"
	"strconv"
)

type CommentController interface {
	PostComment(ctx *gin.Context)
	GetAllComments(ctx *gin.Context)
	PutComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentHttp struct {
	commentUsecase comment.CommentUsecase
}

func NewCommentHttp(commentUsecase comment.CommentUsecase) CommentController {
	return &commentHttp{
		commentUsecase: commentUsecase,
	}
}

func (c commentHttp) PostComment(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint)

	var bindComment domain.Comment

	err := ctx.ShouldBindJSON(&bindComment)
	if err != nil {
		ctx.Abort()

		return
	}

	addedComment, err := c.commentUsecase.AddComment(ctx, &userId, &bindComment)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusCreated, addedComment)
}

func (c commentHttp) GetAllComments(ctx *gin.Context) {
	comments, err := c.commentUsecase.GetAllComments(ctx)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (c commentHttp) PutComment(ctx *gin.Context) {
	paramCommentId := ctx.Param("commentId")
	conv, _ := strconv.ParseUint(paramCommentId, 10, 64)
	commentId := uint(conv)

	userId := ctx.MustGet("userId").(uint)

	var bindComment domain.Comment

	err := ctx.ShouldBindJSON(&bindComment)
	if err != nil {
		ctx.Abort()

		return
	}

	updatedComment, err := c.commentUsecase.UpdateComment(ctx, &commentId, &userId, &bindComment)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, updatedComment)
}

func (c commentHttp) DeleteComment(ctx *gin.Context) {
	paramPhotoId := ctx.Param("commentId")
	conv, _ := strconv.ParseUint(paramPhotoId, 10, 64)
	commentId := uint(conv)

	userId := ctx.MustGet("userId").(uint)

	_, err := c.commentUsecase.DeleteComment(ctx, &commentId, &userId)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	})
}
