package entity

type Address struct {
	ID            int    `gorm:"primary_key" json:"id"`
	Receiver      string `json:"receiver"`
	PhoneReceiver string `json:"phone_receiver"`
	AddressDetail string `json:"address_detail"`
	Province      string `json:"province"`
	City          string `json:"city"`
	UserID        int    `json:"user_id"`
}

type AddressInput struct {
	Receiver      string `json:"receiver" binding:"required"`
	PhoneReceiver string `json:"phone_receiver" binding:"required"`
	AddressDetail string `json:"address_detail" binding:"required"`
	Province      string `json:"province" binding:"required"`
	City          string `json:"city" binding:"required"`
}

type AddressUpdate struct {
	Receiver      string `json:"receiver"`
	PhoneReceiver string `json:"phone_receiver"`
	AddressDetail string `json:"address_detail"`
	Province      string `json:"province"`
	City          string `json:"city"`
}

type AddressResponse struct {
	ID            int    `gorm:"primary_key" json:"-"`
	Receiver      string `json:"receiver"`
	PhoneReceiver string `json:"phone_receiver"`
	AddressDetail string `json:"address_detail"`
	Province      string `json:"province"`
	City          string `json:"city"`
	UserID        int    `json:"-"`
}
