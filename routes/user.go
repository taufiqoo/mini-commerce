package routes

import (
	"mini-commerce/config"
	"mini-commerce/handler"
	"mini-commerce/layer/user"
	"mini-commerce/middleware"

	"github.com/gin-gonic/gin"
)

var (
	DB          = config.Connection()
	userRepo    = user.NewRepository(DB)
	userService = user.NewService(userRepo)
	authService = middleware.NewService()
	userHandler = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.GET("/users", middleware.Authentication(authService), userHandler.GetUserByIdHandler)
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.POST("users/refresh-token", userHandler.RefreshTokenHandler)
	r.POST("/users/logout", middleware.Authentication(authService), userHandler.LogoutHandler)
}
