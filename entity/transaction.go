package entity

import "time"

type Transaction struct {
	ID            int       `gorm:"primary_key" json:"id"`
	Date          time.Time `json:"date"`
	UserID        int       `json:"user_id"`
	AddressID     int       `json:"address_id"`
	CartID        int       `json:"cart_id"`
	ProductID     int       `json:"product_id"`
	ProductName   string    `json:"product_name"`
	Quantity      int       `json:"quantity"`
	TotalPrice    int       `json:"total_price"`
	PaymentMethod string    `json:"payment_method"`
	Status        string    `json:"status" gorm:"default:pending"`
	Address       Address   `json:"address"`
}

type TransactionInput struct {
	AddressID     int    `json:"address_id" binding:"required"`
	CartID        int    `json:"cart_id" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

type PaymentInput struct {
	Nominal float64 `json:"nominal" binding:"required"`
}
