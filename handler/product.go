package handler

import (
	"mini-commerce/entity"
	"mini-commerce/helper"
	"mini-commerce/layer/product"
	"strconv"

	"github.com/gin-gonic/gin"
)

type prodcutHandler struct {
	service product.Service
}

func NewProductHandler(service product.Service) *prodcutHandler {
	return &prodcutHandler{service}
}

func (h *prodcutHandler) GetAllProductHandler(c *gin.Context) {
	product, err := h.service.GetAllProduct()
	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}

	response := helper.APIResponse(200, "success get all product", product)
	c.JSON(200, response)
}

func (h *prodcutHandler) GetProductByIdHandler(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))

	product, err := h.service.GetProductById(productId)
	if err != nil {
		responseErr := helper.APIFailure(404, "Not found", err.Error())
		c.JSON(404, responseErr)
		return
	}

	response := helper.APIResponse(200, "success get product by id", product)
	c.JSON(200, response)
}

func (h *prodcutHandler) CreateProductHandler(c *gin.Context) {
	var inputProduct entity.ProductInput

	if err := c.ShouldBindJSON(&inputProduct); err != nil {
		splitErr := helper.SplitErrorInformation(err)
		responseErr := helper.APIFailure(400, "All field are required", gin.H{"error": splitErr})
		c.JSON(400, responseErr)
		return
	}

	newProduct, err := h.service.SaveNewProduct(inputProduct, c)

	if err != nil {
		responseErr := helper.APIFailure(500, "internal server error", err.Error())
		c.JSON(500, responseErr)
		return
	}
	response := helper.APIResponse(201, "success create product", newProduct)
	c.JSON(201, response)
}

func (h *prodcutHandler) UpdateProductHandler(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))

	var updateProductInput entity.ProductUpdateInput

	if err := c.ShouldBindJSON(&updateProductInput); err != nil {
		responseErr := helper.APIFailure(500, "internal server error", err.Error())

		c.JSON(500, responseErr)
		return
	}

	product, err := h.service.UpdateProduct(productId, updateProductInput)

	if err != nil {
		responseErr := helper.APIFailure(404, "not found", err.Error())
		c.JSON(404, responseErr)
		return
	}

	response := helper.APIResponse(200, "success update product", product)
	c.JSON(200, response)
}

func (h *prodcutHandler) DeleteProductHandler(c *gin.Context) {
	productId, _ := strconv.Atoi(c.Param("productId"))

	product, err := h.service.DeleteProduct(productId)
	if err != nil {
		responseErr := helper.APIFailure(404, "Product not found", err.Error())
		c.JSON(404, responseErr)
		return
	}

	response := helper.FormatDelete(200, product)
	c.JSON(200, response)
}
