package migration

import "time"

type User struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	PhotoProfile string    `json:"photo_profile"`
	Password     string    `json:"password"`
	Role         string    `json:"role"`
	Address      []Address `gorm:"foreignKey:UserID" json:"address"`
	Product      []Product `gorm:"foreignKey:UserID" json:"product"`
	Cart         []Cart    `gorm:"foreignKey:UserID" json:"cart"`
}

type Address struct {
	ID            int    `gorm:"primary_key" json:"id"`
	Receiver      string `json:"receiver"`
	PhoneReceiver string `json:"phone_receiver"`
	AddressDetail string `json:"address_detail"`
	Province      string `json:"province"`
	City          string `json:"city"`
	UserID        int    `json:"user_id"`
}

type Product struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	PhotoProduct string `json:"photo_product"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	UserID       int    `json:"user_id"`
	User         []User `gorm:"many2many:cart;" json:"users"`
	Cart         []Cart `gorm:"foreignKey:ProductID" json:"cart"`
}

type Cart struct {
	ID          int    `gorm:"primary_key" json:"id"`
	UserID      int    `json:"user_id"`
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"total_price"`
}

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
	Status        string    `json:"status"`
}
