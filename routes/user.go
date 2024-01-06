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
	r.POST("/auth/register", userHandler.CreateUserHandler)
	r.POST("/auth/login", userHandler.LoginUserHandler)
	r.POST("auth/refresh-token", userHandler.RefreshTokenHandler)
	r.POST("/auth/logout", middleware.Authentication(authService), userHandler.LogoutHandler)
}
