package entity

type Cart struct {
	ID          int    `gorm:"primary_key" json:"id"`
	UserID      int    `json:"user_id"`
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"total_price"`
}

type CartInput struct {
	Quantity int `json:"quantity"`
}
