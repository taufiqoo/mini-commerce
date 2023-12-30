package handler

import (
	"mini-commerce/entity"
	"mini-commerce/helper"
	"mini-commerce/layer/address"
	"strconv"

	"github.com/gin-gonic/gin"
)

type addressHandler struct {
	service address.Service
}

func NewAddressHandler(service address.Service) *addressHandler {
	return &addressHandler{service}
}

func (h *addressHandler) GetAddressByUserIdHandler(c *gin.Context) {
	userId := c.MustGet("currentUser").(int)

	address, err := h.service.GetAddressByUserId(userId)
	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", gin.H{"message": err.Error()})

		c.JSON(500, responseErr)
	}

	response := helper.APIResponse(200, "success get address by user id", address)
	c.JSON(200, response)
}

func (h *addressHandler) GetAddressByIdHandler(c *gin.Context) {
	addressId := c.Param("addressId")
	addressIdInt, _ := strconv.Atoi(addressId)

	address, err := h.service.GetAddressById(addressIdInt)
	if err != nil {
		responseErr := helper.APIFailure(404, "Not found", gin.H{"message": err.Error()})
		c.JSON(404, responseErr)
		return
	}

	response := helper.APIResponse(200, "success get address by id", address)
	c.JSON(200, response)
}

func (h *addressHandler) CreateAddressHandler(c *gin.Context) {
	userId := c.MustGet("currentUser").(int)
	var inputAddress entity.AddressInput

	if err := c.ShouldBindJSON(&inputAddress); err != nil {
		splitErr := helper.SplitErrorInformation(err)
		responseErr := helper.APIFailure(400, "All field are required", gin.H{"message": splitErr})
		c.JSON(400, responseErr)
		return
	}

	newAddress, err := h.service.SaveNewAddress(userId, inputAddress)
	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", gin.H{"message": err.Error()})
		c.JSON(500, responseErr)
		return
	}
	response := helper.APIResponse(201, "success create new address", newAddress)
	c.JSON(201, response)
}

func (h *addressHandler) UpdateAddressHandler(c *gin.Context) {
	addressId := c.Param("addressId")
	addressIdInt, _ := strconv.Atoi(addressId)

	var updateAddress entity.AddressUpdate
	if err := c.ShouldBindJSON(&updateAddress); err != nil {
		splitErr := helper.SplitErrorInformation(err)
		responseErr := helper.APIFailure(400, "Input data required", gin.H{"message": splitErr})
		c.JSON(400, responseErr)
		return
	}

	address, err := h.service.UpdateAddress(addressIdInt, updateAddress)
	if err != nil {
		responseErr := helper.APIFailure(404, "Not found", gin.H{"message": err.Error()})
		c.JSON(404, responseErr)
		return
	}

	response := helper.APIResponse(200, "success update address", address)
	c.JSON(200, response)
}

func (h *addressHandler) DeleteAddressHandler(c *gin.Context) {
	addressId := c.Param("addressId")
	addressIdInt, _ := strconv.Atoi(addressId)

	address, err := h.service.DeleteAddress(addressIdInt)
	if err != nil {
		responseErr := helper.APIFailure(404, "Not found", gin.H{"message": err.Error()})
		c.JSON(404, responseErr)
		return
	}

	response := helper.APIResponse(200, "success delete address", address)
	c.JSON(200, response)
}
