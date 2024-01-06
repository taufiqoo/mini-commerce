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
	r.GET("/address/:addressId", middleware.Authentication(authService), middleware.Authorization("user"), addressHandler.GetAddressByIdHandler)
	r.POST("/address", middleware.Authentication(authService), middleware.Authorization("user"), addressHandler.CreateAddressHandler)
	r.PUT("/address/:addressId", middleware.Authentication(authService), middleware.Authorization("user"), addressHandler.UpdateAddressHandler)
	r.DELETE("/address/:addressId", middleware.Authentication(authService), middleware.Authorization("user"), addressHandler.DeleteAddressHandler)
}
