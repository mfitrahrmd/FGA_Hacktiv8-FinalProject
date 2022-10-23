package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf(e.Message)
}

func NewError(statusCode int, message string) Error {
	return Error{
		StatusCode: statusCode,
		Message:    message,
	}
}

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()
}
