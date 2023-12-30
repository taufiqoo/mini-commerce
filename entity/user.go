package entity

type User struct {
	ID           int       `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	Email        string    `gorm:"unique, email" json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	PhotoProfile string    `json:"photo_profile"`
	Password     string    `gorm:"password" json:"-"`
	Role         string    `json:"role" gorm:"default:user"`
	Products     []Product `json:"products"`
	Address      []Address `json:"address"`
}

type UserInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required" gorm:"unique, email"`
	PhoneNumber string `json:"phone_number" binding:"required" gorm:"unique"`
	Password    string `json:"password" binding:"required" gorm:"password"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
