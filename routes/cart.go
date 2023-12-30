package routes

import (
	"mini-commerce/handler"
	"mini-commerce/layer/cart"
	"mini-commerce/middleware"

	"github.com/gin-gonic/gin"
)

var (
	cartRepo    = cart.NewRepository(DB)
	cartService = cart.NewService(cartRepo)
	cartHandler = handler.NewCartHandler(cartService)
)

func CartRoute(r *gin.Engine) {
	r.GET("/carts", middleware.Authentication(authService), middleware.Authorization("user"), cartHandler.GetMyCartHandler)
	r.POST("/carts/:productId", middleware.Authentication(authService), middleware.Authorization("user"), cartHandler.CreateCartHandler)
	r.PUT("/carts/:cartId/product/:productId", middleware.Authentication(authService), middleware.Authorization("user"), cartHandler.UpdateCartHandler)
	r.DELETE("/carts/:cartId", middleware.Authentication(authService), middleware.Authorization("user"), cartHandler.DeleteCartHandler)
}
