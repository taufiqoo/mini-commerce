package handler

import (
	"mini-commerce/entity"
	"mini-commerce/helper"
	"mini-commerce/layer/cart"
	"strconv"

	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	service cart.Service
}

func NewCartHandler(service cart.Service) *cartHandler {
	return &cartHandler{service}
}

func (h *cartHandler) GetMyCartHandler(c *gin.Context) {
	userID, exists := c.Get("currentUser")
	if !exists {
		responseErr := helper.APIFailure(400, "bad request", "user ID not found in context")
		c.JSON(400, responseErr)
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		responseErr := helper.APIFailure(400, "bad request", "invalid user ID type in context")
		c.JSON(400, responseErr)
		return
	}

	cart, err := h.service.GetMyCart(userIDInt)
	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}

	response := helper.APIResponse(200, "success get my cart", cart)
	c.JSON(200, response)
}

func (h *cartHandler) CreateCartHandler(c *gin.Context) {
	var cartInput entity.CartInput
	if err := c.ShouldBindJSON(&cartInput); err != nil {
		responseErr := helper.APIFailure(400, "bad request", "invalid request body")
		c.JSON(400, responseErr)
		return
	}

	cart := entity.CartInput{
		Quantity: cartInput.Quantity,
	}

	newCart, err := h.service.SaveNewCart(cart, c)
	if err != nil {
		responseErr := helper.APIFailure(400, "Bad Request", err.Error())
		c.JSON(400, responseErr)
		return
	}

	response := helper.APIResponse(201, "success create cart", newCart)
	c.JSON(201, response)
}

func (h *cartHandler) UpdateCartHandler(c *gin.Context) {
	var cartInput entity.CartInput
	if err := c.ShouldBindJSON(&cartInput); err != nil {
		responseErr := helper.APIFailure(400, "bad request", "invalid request body")
		c.JSON(400, responseErr)
		return
	}

	updatedCart, err := h.service.UpdateCart(cartInput, c)
	if err != nil {
		responseErr := helper.APIFailure(404, "product or cart not found", err.Error())
		c.JSON(404, responseErr)
		return
	}

	response := helper.APIResponse(200, "success update cart", updatedCart)
	c.JSON(200, response)
}

func (h *cartHandler) DeleteCartHandler(c *gin.Context) {
	cartID, err := strconv.Atoi(c.Param("cartId"))
	if err != nil {
		responseErr := helper.APIFailure(400, "bad request", "invalid cart ID")
		c.JSON(400, responseErr)
		return
	}

	deletedCart, err := h.service.DeleteCart(cartID)
	if err != nil {
		responseErr := helper.APIFailure(404, "Cart not found", err.Error())
		c.JSON(404, responseErr)
		return
	}

	response := helper.APIResponse(200, "success delete cart", deletedCart)
	c.JSON(200, response)
}
