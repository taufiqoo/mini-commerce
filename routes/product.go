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
	r.GET("/product/:productId", productHandler.GetProductByIdHandler)

	r.POST("/product", middleware.Authentication(authService), middleware.Authorization("admin"), productHandler.CreateProductHandler)
	r.PUT("/product/:productId", middleware.Authentication(authService), middleware.Authorization("admin"), productHandler.UpdateProductHandler)
	r.DELETE("/product/:productId", middleware.Authentication(authService), middleware.Authorization("admin"), productHandler.DeleteProductHandler)
}
