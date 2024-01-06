package transaction

import (
	"mini-commerce/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAllTransaction(userId int) ([]entity.Transaction, error)
	FindAllUserTransaction() ([]entity.Transaction, error)
	FindTransactionById(id int) (entity.Transaction, error)
	CreateTransaction(transaction entity.Transaction) (entity.Transaction, error)
	DeleteTransaction(id int) (interface{}, error)
	FindCartById(cartId int) (entity.Cart, error)
	UpdateStatusTransaction(transaction entity.Transaction) (entity.Transaction, error)
	UpdateProductQuantity(product entity.Product) (entity.Product, error)
	FindProductById(productId int) (entity.Product, error)
	DeleteCart(cartId int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllTransaction(userId int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if err := r.db.Where("user_id = ?", userId).Preload("Address").Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) FindAllUserTransaction() ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if err := r.db.Preload("Address").Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) FindTransactionById(id int) (entity.Transaction, error) {
	var transaction entity.Transaction

	if err := r.db.Where("id = ?", id).Preload("Address").First(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) CreateTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	if err := r.db.Preload("Address").Create(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) DeleteTransaction(id int) (interface{}, error) {
	if err := r.db.Where("id = ?", id).Delete(&entity.Transaction{}).Error; err != nil {
		return "error", err
	}

	status := "transaction successfully deleted"
	return status, nil
}

func (r *repository) FindCartById(cartId int) (entity.Cart, error) {
	var cart entity.Cart

	if err := r.db.Where("id = ?", cartId).First(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

func (r *repository) FindProductById(productId int) (entity.Product, error) {
	var product entity.Product

	if err := r.db.Where("id = ?", productId).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) UpdateStatusTransaction(transaction entity.Transaction) (entity.Transaction, error) {
	if err := r.db.Save(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) UpdateProductQuantity(product entity.Product) (entity.Product, error) {
	if err := r.db.Save(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) DeleteCart(cartId int) error {
	if err := r.db.Where("id = ?", cartId).Delete(&entity.Cart{}).Error; err != nil {
		return err
	}
	return nil
}
