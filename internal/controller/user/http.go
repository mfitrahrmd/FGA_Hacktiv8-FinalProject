package user

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/user"
)

type UserController interface {
	PostUserRegister(ctx *gin.Context)
	PostUserLogin(ctx *gin.Context)
	PutUserUpdateData(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userHttp struct {
	userUsecase user.UserUsecase
	ctx         context.Context
}

func NewUserHttp(ctx context.Context, userUsecase user.UserUsecase) UserController {
	return &userHttp{
		userUsecase: userUsecase,
		ctx:         ctx,
	}
}

func (u userHttp) PostUserRegister(ctx *gin.Context) {
	var bindUser domain.UserRegister

	err := ctx.ShouldBindJSON(&bindUser)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	registeredUser, err := u.userUsecase.Register(u.ctx, &domain.User{
		Username: bindUser.Username,
		Email:    bindUser.Email,
		Password: bindUser.Password,
	})
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	registeredUser.Password = ""

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (u userHttp) PostUserLogin(ctx *gin.Context) {
	var bindUser domain.UserLogin

	err := ctx.ShouldBindJSON(&bindUser)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	token, err := u.userUsecase.Login(u.ctx, &domain.User{
		Email:    bindUser.Email,
		Password: bindUser.Password,
	})
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (u userHttp) PutUserUpdateData(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint)

	var bindUser domain.UserUpdateData

	err := ctx.ShouldBindJSON(&bindUser)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind)

		return
	}

	updatedUser, err := u.userUsecase.UpdateData(u.ctx, &userId, &domain.User{
		Email:    bindUser.Email,
		Username: bindUser.Username,
	})
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	updatedUser.Password = ""

	ctx.JSON(http.StatusOK, updatedUser)
}

func (u userHttp) DeleteUser(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint)

	_, err := u.userUsecase.Delete(u.ctx, &userId)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePublic)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
