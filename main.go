package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/domain"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/db"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/config/env"
	commentController "github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/controller/comment"
	"github.com/mfitrahrmd420/FGA_Hacktiv8-FinalProject/internal/controller/middleware"
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
)

func main() {
	env.LoadEnvFile()

	conn := db.GetPostgresGorm()
	conn.Debug().AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.SocialMedia{})

	ur := userRepository.NewUserRepository(conn)
	uu := userUsecase.NewUserUsecase(ur)
	pr := photoRepository.NewPhotoRepository(conn)
	pu := photoUsecase.NewPhotoUsecase(pr)
	cr := commentRepository.NewCommentRepository(conn)
	cu := commentUsecase.NewCommentUsecase(cr, ur, pr)
	smr := socialMediaRepository.NewSocialMediaRepository(conn)
	smu := socialMediaUsecase.NewSocialMediaUsecase(smr, pr)

	r := gin.Default()

	r.Use(middleware.ErrorHandler)

	uHttp := userController.NewUserHttp(uu)
	pHttp := photoController.NewUserHttp(pu)
	cHttp := commentController.NewCommentHttp(cu)
	smHttp := socialMediaController.NewSocialMediaController(smu)

	rUsers := r.Group("/users")
	rUsers.POST("/register", uHttp.PostUserRegister)
	rUsers.POST("/login", uHttp.PostUserLogin)

	rUsers.Use(middleware.Authentication)

	rUsers.PUT("", uHttp.PutUserUpdateData)
	rUsers.DELETE("", uHttp.DeleteUser)

	rPhoto := r.Group("/photos")
	rPhoto.Use(middleware.Authentication)
	rPhoto.POST("", pHttp.PostPhoto)
	rPhoto.GET("", pHttp.GetAllPhotos)
	rPhoto.PUT("/:photoId", pHttp.PutPhoto)
	rPhoto.DELETE("/:photoId", pHttp.DeletePhoto)

	rComment := r.Group("/comments")
	rComment.Use(middleware.Authentication)
	rComment.POST("", cHttp.PostComment)
	rComment.GET("", cHttp.GetAllComments)
	rComment.PUT("/:commentId", cHttp.PutComment)
	rComment.DELETE("/:commentId", cHttp.DeleteComment)

	rSocialMedia := r.Group("socialmedias")
	rSocialMedia.Use(middleware.Authentication)
	rSocialMedia.POST("", smHttp.PostSocialMedia)
	rSocialMedia.GET("", smHttp.GetAllSocialMedias)
	rSocialMedia.PUT("/:socialMediaId", smHttp.PutSocialMedia)
	rSocialMedia.DELETE("/:socialMediaId", smHttp.DeleteSocialMedia)

	r.Run(":8001")
}
