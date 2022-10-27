package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	docs "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/docs"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/db"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
	commentController "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/controller/comment"
	middlewareController "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/controller/middleware"
	photoController "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/controller/photo"
	socialMediaController "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/controller/socialMedia"
	userController "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/controller/user"
	commentRepository "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/repository/comment"
	photoRepository "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/repository/photo"
	socialMediaRepository "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/repository/socialMedia"
	userRepository "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/repository/user"
	commentUsecase "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/comment"
	photoUsecase "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/photo"
	socialMediaUsecase "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/socialMedia"
	userUsecase "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/service/user"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Photos & Social Meida Management API
// @version 1.0
// @description This is a sample gin server for managing user's photos and social medias with Authentication & Authorization.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /
func main() {
	// load env file
	env.LoadEnvFile()

	ctx := context.Background()

	// setup database
	conn := db.GetPostgresGorm()
	conn.Debug().AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.SocialMedia{})

	// bootstraping api
	ur := userRepository.NewUserRepository(conn)
	uu := userUsecase.NewUserUsecase(ur)
	uHttp := userController.NewUserHttp(ctx, uu)
	pr := photoRepository.NewPhotoRepository(conn)
	pu := photoUsecase.NewPhotoUsecase(pr)
	pHttp := photoController.NewUserHttp(ctx, pu)
	cr := commentRepository.NewCommentRepository(conn)
	cu := commentUsecase.NewCommentUsecase(cr, ur, pr)
	cHttp := commentController.NewCommentHttp(ctx, cu)
	smr := socialMediaRepository.NewSocialMediaRepository(conn)
	smu := socialMediaUsecase.NewSocialMediaUsecase(smr, pr)
	smHttp := socialMediaController.NewSocialMediaController(ctx, smu)

	r := gin.Default()

	// swagger API documentation route
	// http://127.0.0.1:8001/swagger/index.html
	docs.SwaggerInfo.BasePath = "/"
	ginSwagger.DefaultModelsExpandDepth(1)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// add custom tag struct validation
	middlewareController.RegisterCustomValidation(ctx, ur)

	// gin cors
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:    []string{"*"},
	}))

	// passing all errors to error middleware
	r.Use(middlewareController.ErrorHandler)

	rUsers := r.Group("/users")
	{
		rUsers.POST("/register", uHttp.PostUserRegister)
		rUsers.POST("/login", uHttp.PostUserLogin)

		rUsers.Use(middlewareController.Authentication)
		{
			rUsers.PUT("", uHttp.PutUserUpdateData)
			rUsers.DELETE("", uHttp.DeleteUser)
		}
	}

	rPhoto := r.Group("/photos")
	{
		rPhoto.Use(middlewareController.Authentication)
		{
			rPhoto.POST("", pHttp.PostPhoto)
			rPhoto.GET("", pHttp.GetAllPhotos)
			rPhoto.PUT("/:photoId", pHttp.PutPhoto)
			rPhoto.DELETE("/:photoId", pHttp.DeletePhoto)
		}
	}

	rComment := r.Group("/comments")
	{
		rComment.Use(middlewareController.Authentication)
		{
			rComment.POST("", cHttp.PostComment)
			rComment.GET("", cHttp.GetAllComments)
			rComment.PUT("/:commentId", cHttp.PutComment)
			rComment.DELETE("/:commentId", cHttp.DeleteComment)
		}
	}

	rSocialMedia := r.Group("socialmedias")
	{
		rSocialMedia.Use(middlewareController.Authentication)
		{
			rSocialMedia.POST("", smHttp.PostSocialMedia)
			rSocialMedia.GET("", smHttp.GetAllSocialMedias)
			rSocialMedia.PUT("/:socialMediaId", smHttp.PutSocialMedia)
			rSocialMedia.DELETE("/:socialMediaId", smHttp.DeleteSocialMedia)
		}
	}

	log.Fatalln(r.Run(fmt.Sprintf("%s:%s", env.HOST, env.PORT)))
}
