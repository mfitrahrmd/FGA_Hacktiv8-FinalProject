package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/user"
	"net/http"
)

type UserController interface {
	PostUserRegister(ctx *gin.Context)
	PostUserLogin(ctx *gin.Context)
	PutUserUpdateData(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userHttp struct {
	userUsecase user.UserUsecase
}

func NewUserHttp(userUsecase user.UserUsecase) UserController {
	return &userHttp{
		userUsecase: userUsecase,
	}
}

func (u userHttp) PostUserRegister(ctx *gin.Context) {
	var user domain.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.Abort()

		return
	}

	registeredUser, err := u.userUsecase.Register(ctx, &user)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (u userHttp) PostUserLogin(ctx *gin.Context) {
	var bindUser domain.User

	err := ctx.ShouldBindJSON(&bindUser)
	if err != nil {
		ctx.Abort()

		return
	}

	token, err := u.userUsecase.Login(ctx, &bindUser)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (u userHttp) PutUserUpdateData(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint)

	var bindUser domain.User

	err := ctx.ShouldBindJSON(&bindUser)
	if err != nil {
		ctx.Abort()

		return
	}

	updatedUser, err := u.userUsecase.UpdateData(ctx, &userId, &bindUser)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (u userHttp) DeleteUser(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint)

	_, err := u.userUsecase.Delete(ctx, &userId)
	if err != nil {
		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	})
}
