package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service"
)

type Error struct {
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf(e.Message)
}

func NewError(statusCode int, message string) Error {
	return Error{
		Message: message,
	}
}

func customValidationError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required.", err.Field())
	case "min":
		switch err.Kind().String() {
		case "string":
			return fmt.Sprintf("%s must be longer than or equal %s characters.", err.Field(), err.Param())
		case "int":
			return fmt.Sprintf("%s must be greater than or equal to %s.", err.Field(), err.Param())
		default:
			return err.Error()
		}
	case "max":
		switch err.Kind().String() {
		case "string":
			return fmt.Sprintf("%s cannot be longer than %s characters.", err.Field(), err.Param())
		case "int":
			return fmt.Sprintf("%s cannot greater than %s.", err.Field(), err.Param())
		default:
			return err.Error()
		}
	case "email":
		return fmt.Sprintf("%s should be valid email address", err.Field())
	case "uniqueEmail":
		return fmt.Sprintf("%s already in used", err.Field())
	case "uniqueUsername":
		return fmt.Sprintf("%s already in used", err.Field())
	default:
		return err.Error()
	}
}

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		for _, err := range ctx.Errors {
			switch err.Type {
			case gin.ErrorTypeBind:
				errMap := make(map[string]string)
				errs, ok := err.Err.(validator.ValidationErrors)
				if !ok {
					ctx.JSON(http.StatusBadRequest, Error{
						Message: "request body should be valid JSON",
					})

					return
				}
				for _, fieldErr := range []validator.FieldError(errs) {
					errMap[fieldErr.Field()] = customValidationError(fieldErr)
				}

				ctx.JSON(http.StatusBadRequest, Error{
					Message: func() string {
						var msg string
						for _, e := range errMap {
							if msg == "" {
								msg = e
								continue
							}
							msg = fmt.Sprintf("%s %s", msg, e)
						}

						return msg
					}(),
				})

				return
			case gin.ErrorTypePublic:
				var serviceError service.ServiceError
				if errors.As(err, &serviceError) {
					ctx.JSON(serviceError.StatusCode, Error{
						Message: serviceError.Error(),
					})

					return
				}
			default:
				log.Println(err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"message": "unexpected server error",
				})

				return
			}
		}
	}
}
