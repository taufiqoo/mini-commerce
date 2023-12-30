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
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllTransaction(userId int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if err := r.db.Where("user_id = ?", userId).Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) FindAllUserTransaction() ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	if err := r.db.Find(&transactions).Error; err != nil {
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
	if err := r.db.Create(&transaction).Error; err != nil {
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
