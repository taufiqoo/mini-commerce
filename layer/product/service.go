package product

import (
	"errors"
	"fmt"
	"mini-commerce/entity"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAllProduct() ([]entity.Product, error)
	GetProductById(id int) (entity.Product, error)
	SaveNewProduct(input entity.ProductInput, c *gin.Context) (entity.Product, error)
	DeleteProduct(id int) (interface{}, error)
	UpdateProduct(id int, dataUpdate entity.ProductUpdateInput) (entity.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllProduct() ([]entity.Product, error) {
	products, err := s.repository.FindAllProduct()
	if err != nil {
		return products, err
	}
	return products, nil
}

func (s *service) GetProductById(id int) (entity.Product, error) {
	product, err := s.repository.FindProductById(id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (s *service) SaveNewProduct(input entity.ProductInput, c *gin.Context) (entity.Product, error) {
	userID, exists := c.Get("currentUser")
	if !exists {
		return entity.Product{}, errors.New("user ID not found in context")
	}

	userIDInt, ok := userID.(int)
	if !ok {
		return entity.Product{}, errors.New("invalid user ID type in context")
	}

	product := entity.Product{
		Name:         input.Name,
		Description:  input.Description,
		PhotoProduct: input.PhotoProduct,
		Price:        input.Price,
		Stock:        input.Stock,
		UserID:       userIDInt,
	}
	createProduct, err := s.repository.CreateProduct(product)
	if err != nil {
		return createProduct, err
	}
	return createProduct, nil
}

func (s *service) DeleteProduct(id int) (interface{}, error) {
	product, err := s.repository.FindProductById(id)
	if err != nil {
		return nil, err
	}
	if product.ID == 0 {
		newError := fmt.Sprintf("product with id %d not found", id)
		return nil, errors.New(newError)
	}

	deleteProduct, err := s.repository.DeleteProduct(id)
	if err != nil {
		return deleteProduct, err
	}
	return deleteProduct, nil
}

func (s *service) UpdateProduct(id int, dataInput entity.ProductUpdateInput) (entity.Product, error) {
	var dataUpdate = map[string]interface{}{}

	product, err := s.repository.FindProductById(id)
	if err != nil {
		return entity.Product{}, err
	}
	if product.ID == 0 {
		return entity.Product{}, errors.New("product id not found")
	}

	if dataInput.Name != "" || len(dataInput.Name) != 0 {
		dataUpdate["name"] = dataInput.Name
	}
	if dataInput.Description != "" || len(dataInput.Description) != 0 {
		dataUpdate["description"] = dataInput.Description
	}
	if dataInput.PhotoProduct != "" || len(dataInput.PhotoProduct) != 0 {
		dataUpdate["photo_product"] = dataInput.PhotoProduct
	}
	if dataInput.Price != 0 {
		dataUpdate["price"] = dataInput.Price
	}
	if dataInput.Stock != 0 {
		dataUpdate["stock"] = dataInput.Stock
	}

	productUpdated, err := s.repository.UpdateProduct(id, dataUpdate)
	if err != nil {
		return entity.Product{}, err
	}
	return productUpdated, nil
}
