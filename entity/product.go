package entity

type Product struct {
	ID           int    `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	PhotoProduct string `json:"photo_product"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	UserID       int    `json:"user_id"`
}

type ProductInput struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description" binding:"required"`
	PhotoProduct string `json:"photo_product" binding:"required"`
	Price        int    `json:"price" binding:"required"`
	Stock        int    `json:"stock" binding:"required"`
	UserID       int    `json:"user_id"`
}

type ProductUpdateInput struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	PhotoProduct string `json:"photo_product"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
}
