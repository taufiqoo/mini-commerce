package cart

import (
	"errors"
	"mini-commerce/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetMyCart(userId int) ([]entity.Cart, error)
	SaveNewCart(input entity.CartInput, c *gin.Context) (entity.Cart, error)
	UpdateCart(inputUpdate entity.CartInput, c *gin.Context) (entity.Cart, error)
	DeleteCart(cartId int) (interface{}, error)
	FindProductById(id int) (entity.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetMyCart(userId int) ([]entity.Cart, error) {
	carts, err := s.repository.FindMyCart(userId)
	if err != nil {
		return carts, err
	}
	return carts, nil
}

func (s *service) SaveNewCart(input entity.CartInput, c *gin.Context) (entity.Cart, error) {
	userID, exists := c.Get("currentUser")
	if !exists {
		return entity.Cart{}, errors.New("user ID not found in context")
	}

	userIDInt, ok := userID.(int)
	if !ok {
		return entity.Cart{}, errors.New("invalid user ID type in context")
	}

	productId, _ := strconv.Atoi(c.Param("productId"))
	product, err := s.repository.FindProductById(productId)
	if err != nil {
		return entity.Cart{}, err
	}

	if input.Quantity > product.Stock {
		return entity.Cart{}, errors.New("quantity is greater than stock")
	}

	cart := entity.Cart{
		ProductID:   productId,
		ProductName: product.Name,
		Quantity:    input.Quantity,
		TotalPrice:  product.Price * input.Quantity,
		UserID:      userIDInt,
	}

	createCart, err := s.repository.CreateCart(cart)
	if err != nil {
		return entity.Cart{}, err
	}
	return createCart, nil
}

func (s *service) UpdateCart(inputUpdate entity.CartInput, c *gin.Context) (entity.Cart, error) {
	cartId, _ := strconv.Atoi(c.Param("cartId"))
	cart, err := s.repository.FindCartById(cartId)
	if err != nil {
		return entity.Cart{}, err
	}

	productId, _ := strconv.Atoi(c.Param("productId"))
	product, err := s.repository.FindProductById(productId)
	if err != nil {
		return entity.Cart{}, err
	}

	if inputUpdate.Quantity > product.Stock {
		return entity.Cart{}, errors.New("quantity is greater than stock")
	}

	if inputUpdate.Quantity > 0 {
		cart.Quantity = inputUpdate.Quantity
		cart.TotalPrice = cart.Quantity * product.Price
	}

	updatedCart, err := s.repository.UpdateCart(cart)
	if err != nil {
		return entity.Cart{}, err
	}
	return updatedCart, nil
}

func (s *service) DeleteCart(cartId int) (interface{}, error) {
	cart, err := s.repository.DeleteCart(cartId)
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (s *service) FindProductById(id int) (entity.Product, error) {
	product, err := s.repository.FindProductById(id)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}
