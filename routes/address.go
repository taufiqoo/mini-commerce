package routes

import (
	"mini-commerce/handler"
	"mini-commerce/layer/address"
	"mini-commerce/middleware"

	"github.com/gin-gonic/gin"
)

var (
	addressRepo    = address.NewRepository(DB)
	addressService = address.NewService(addressRepo)
	addressHandler = handler.NewAddressHandler(addressService)
)

func AddressRoute(r *gin.Engine) {
	r.GET("/addresses", middleware.Authentication(authService), middleware.Authorization("user"), addressHandler.GetAddressByUserIdHandler)
	r.GET("/addresses/:addressId", middleware.Authentication(authService), middleware.Authorization("user"), addressHandler.GetAddressByIdHandler)
	r.POST("/addresses", middleware.Authentication(authService), middleware.Authorization("user"), addressHandler.CreateAddressHandler)
	r.PUT("/addresses/:addressId", middleware.Authentication(authService), middleware.Authorization("user"), addressHandler.UpdateAddressHandler)
	r.DELETE("/addresses/:addressId", middleware.Authentication(authService), middleware.Authorization("user"), addressHandler.DeleteAddressHandler)
}
