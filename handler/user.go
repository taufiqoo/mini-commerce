package handler

import (
	"context"
	"fmt"
	"mini-commerce/config"
	"mini-commerce/entity"
	"mini-commerce/helper"
	"mini-commerce/layer/user"
	"mini-commerce/middleware"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service     user.Service
	authService middleware.Service
}

func NewUserHandler(service user.Service, authService middleware.Service) *userHandler {
	return &userHandler{service, authService}
}

func (h *userHandler) GetUserByIdHandler(c *gin.Context) {
	userId := int(c.MustGet("currentUser").(int))

	user, err := h.service.GetUserById(userId)
	if err != nil {
		responseError := helper.APIFailure(401, "Unauthorized", gin.H{"message": err.Error()})

		c.JSON(401, responseError)
		return
	}

	response := helper.APIResponse(200, "Success get user by id", user)
	c.JSON(200, response)
}

func (h *userHandler) CreateUserHandler(c *gin.Context) {
	var userInput entity.UserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIFailure(400, "All field are required", gin.H{"message": splitError})

		c.JSON(400, responseError)
		return
	}

	existingEmail, _ := h.service.GetUserByEmail(userInput.Email)
	if existingEmail.ID != 0 {
		responseError := helper.APIFailure(400, "Bad Request", gin.H{"message": "Email already exists"})
		c.JSON(400, responseError)
		return
	}

	newUser, err := h.service.SaveNewUser(userInput)

	if err != nil {
		responseError := helper.APIFailure(500, "Internal server error", gin.H{"message": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse(201, "Success create new user", newUser)
	c.JSON(201, response)
}

func (h *userHandler) LoginUserHandler(c *gin.Context) {
	var userLogin entity.UserLogin

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		splitError := helper.SplitErrorInformation(err)
		responseError := helper.APIFailure(400, "All fields are required", gin.H{"message": splitError})
		c.JSON(400, responseError)
		return
	}

	user, err := h.service.LoginUser(userLogin)
	if err != nil {
		responseError := helper.APIFailure(400, "Invalid email or password", gin.H{"message": err.Error()})
		c.JSON(400, responseError)
		return
	}

	accessToken, err := h.authService.GenerateToken(user.ID, user.Role)
	if err != nil {
		responseError := helper.APIFailure(500, "Internal server error", gin.H{"message": err.Error()})
		c.JSON(500, responseError)
		return
	}

	refreshToken, err := h.authService.GenerateRefreshToken(user.ID, user.Role)
	if err != nil {
		responseError := helper.APIFailure(500, "Internal server error", gin.H{"message": err.Error()})
		c.JSON(500, responseError)
		return
	}

	// Simpan refresh token ke Redis
	redisClient := config.Redis()
	err = redisClient.Set(context.Background(), fmt.Sprintf("refresh_token:%d", user.ID), refreshToken, 7*24*time.Hour).Err()
	if err != nil {
		responseError := helper.APIFailure(500, "Internal server error", gin.H{"message": err.Error()})
		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse(200, "Success login", gin.H{"user": user, "access_token": accessToken, "refresh_token": refreshToken})
	c.JSON(200, response)
}

func (h *userHandler) RefreshTokenHandler(c *gin.Context) {
	refreshToken := c.PostForm("refresh_token")

	token, err := h.authService.ValidateToken(refreshToken)
	if err != nil {
		responseError := helper.APIFailure(400, "Bad request", gin.H{"message": "Invalid refresh token"})
		c.JSON(400, responseError)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		responseError := helper.APIFailure(401, "Unauthorized", gin.H{"message": "Invalid token claims"})
		c.JSON(401, responseError)
		return
	}

	userID := int(claims["user_id"].(float64))

	redisClient := config.Redis()
	storedRefreshToken, err := redisClient.Get(context.Background(), fmt.Sprintf("refresh_token:%d", userID)).Result()
	if err != nil || storedRefreshToken != refreshToken {
		responseError := helper.APIFailure(401, "Unauthorized", gin.H{"message": "Refresh token not found or invalid"})
		c.JSON(401, responseError)
		return
	}

	// Generate access token baru
	newAccessToken, err := h.authService.GenerateToken(userID, "user")
	if err != nil {
		responseError := helper.APIFailure(500, "Internal server error", gin.H{"message": err.Error()})
		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse(200, "Token refreshed", gin.H{"access_token": newAccessToken})
	c.JSON(200, response)
}

func (h *userHandler) LogoutHandler(c *gin.Context) {
	userID := int(c.MustGet("currentUser").(int))

	redisClient := config.Redis()
	err := redisClient.Del(context.Background(), fmt.Sprintf("refresh_token:%d", userID)).Err()
	if err != nil {
		responseError := helper.APIFailure(500, "Internal server error", gin.H{"message": err.Error()})
		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse(200, "Successfully logged out", gin.H{})
	c.JSON(200, response)
}
