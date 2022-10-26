package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/helper/helper_jwt"
)

func Authentication(ctx *gin.Context) {
	tokenHeader := ctx.Request.Header.Get("Authorization")

	if tokenHeader == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "login required",
		})

		return
	}

	token, err := helper_jwt.ValidateToken(env.JWT_KEY, tokenHeader)
	if err != nil {
		if errors.Is(err, jwt.ErrInvalidKey) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid access token",
			})

			return
		}
		if strings.Contains(err.Error(), "token is expired") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token has expired",
			})

			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "unexpected server error",
		})

		return
	}

	ctx.Set("userId", token.Id)
	ctx.Set("userUsername", token.Username)
	ctx.Set("userEmail", token.Email)

	ctx.Next()
}
