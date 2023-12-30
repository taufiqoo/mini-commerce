package transaction

import (
	"errors"
	"fmt"
	"mini-commerce/entity"
	"time"
)

type Service interface {
	SaveNewTransaction(userId int, input entity.TransactionInput) (entity.Transaction, error)
	GetAllTransaction(userId int) ([]entity.Transaction, error)
	GetAllUserTransaction() ([]entity.Transaction, error)
	GetTransactionDetail(transactionId int) (entity.Transaction, error)
	DeleteTransaction(transactionId int) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveNewTransaction(userId int, input entity.TransactionInput) (entity.Transaction, error) {
	cart, err := s.repository.FindCartById(input.CartID)
	if err != nil {
		return entity.Transaction{}, err
	}

	transaction := entity.Transaction{
		UserID:        userId,
		Date:          time.Now(),
		AddressID:     input.AddressID,
		CartID:        input.CartID,
		ProductID:     cart.ProductID,
		ProductName:   cart.ProductName,
		Quantity:      cart.Quantity,
		TotalPrice:    cart.TotalPrice,
		PaymentMethod: input.PaymentMethod,
	}

	newTransaction, err := s.repository.CreateTransaction(transaction)
	if err != nil {
		return entity.Transaction{}, err
	}
	return newTransaction, nil
}

func (s *service) GetAllTransaction(userId int) ([]entity.Transaction, error) {
	transactions, err := s.repository.FindAllTransaction(userId)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetAllUserTransaction() ([]entity.Transaction, error) {
	transactions, err := s.repository.FindAllUserTransaction()
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (s *service) GetTransactionDetail(transactionId int) (entity.Transaction, error) {
	transaction, err := s.repository.FindTransactionById(transactionId)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) DeleteTransaction(transactionId int) (interface{}, error) {
	transaction, err := s.repository.FindTransactionById(transactionId)
	if err != nil {
		return nil, err
	}

	if transaction.ID == 0 {
		newError := fmt.Sprintf("transaction with id %d not found", transactionId)
		return nil, errors.New(newError)
	}

	deleteTransaction, err := s.repository.DeleteTransaction(transactionId)
	if err != nil {
		return nil, err
	}
	return deleteTransaction, nil
}
