package cart

import (
	"mini-commerce/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindMyCart(userId int) ([]entity.Cart, error)
	CreateCart(cart entity.Cart) (entity.Cart, error)
	UpdateCart(cart entity.Cart) (entity.Cart, error)
	DeleteCart(cartId int) (interface{}, error)
	FindCartById(cartId int) (entity.Cart, error)
	FindProductById(id int) (entity.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindMyCart(userId int) ([]entity.Cart, error) {
	var cart []entity.Cart

	if err := r.db.Where("user_id = ?", userId).Find(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository) CreateCart(cart entity.Cart) (entity.Cart, error) {
	if err := r.db.Create(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository) UpdateCart(cart entity.Cart) (entity.Cart, error) {
	if err := r.db.Save(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository) DeleteCart(cartId int) (interface{}, error) {
	if err := r.db.Where("id = ?", cartId).Delete(&entity.Cart{}).Error; err != nil {
		return "error", err
	}
	status := "cart successfully deleted"

	return status, nil
}

func (r *repository) FindCartById(cartId int) (entity.Cart, error) {
	var cart entity.Cart

	if err := r.db.Where("id = ?", cartId).First(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository) FindProductById(id int) (entity.Product, error) {
	var product entity.Product

	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}
