package routes

import (
	"mini-commerce/handler"
	"mini-commerce/layer/product"
	"mini-commerce/middleware"

	"github.com/gin-gonic/gin"
)

var (
	productRepo    = product.NewRepository(DB)
	productService = product.NewService(productRepo)
	productHandler = handler.NewProductHandler(productService)
)

func ProductRoute(r *gin.Engine) {
	r.GET("/products", productHandler.GetAllProductHandler)
	r.GET("/products/:productId", productHandler.GetProductByIdHandler)

	r.POST("/products", middleware.Authentication(authService), middleware.Authorization("admin"), productHandler.CreateProductHandler)
	r.PUT("/products/:productId", middleware.Authentication(authService), middleware.Authorization("admin"), productHandler.UpdateProductHandler)
	r.DELETE("/products/:productId", middleware.Authentication(authService), middleware.Authorization("admin"), productHandler.DeleteProductHandler)
}
