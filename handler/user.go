package handler

import (
	"mini-commerce/entity"
	"mini-commerce/helper"
	"mini-commerce/layer/user"
	"mini-commerce/middleware"

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
		responseError := helper.APIFailure(400, "All field are required", gin.H{"message": splitError})

		c.JSON(400, responseError)
		return
	}

	user, err := h.service.LoginUser(userLogin)

	if err != nil {
		responseError := helper.APIFailure(400, "Invalid email or password", gin.H{"message": err.Error()})

		c.JSON(400, responseError)
		return
	}

	token, err := h.authService.GenerateToken(user.ID, user.Role)
	if err != nil {
		responseError := helper.APIFailure(500, "Internal server error", gin.H{"message": err.Error()})

		c.JSON(500, responseError)
		return
	}

	response := helper.APIResponse(200, "Success login", gin.H{"user": user, "Authorization": token})
	c.JSON(200, response)
}
