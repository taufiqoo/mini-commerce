package product

import (
	"mini-commerce/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAllProduct() ([]entity.Product, error)
	FindProductById(id int) (entity.Product, error)
	CreateProduct(product entity.Product) (entity.Product, error)
	UpdateProduct(id int, dataUpdate map[string]interface{}) (entity.Product, error)
	DeleteProduct(id int) (interface{}, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllProduct() ([]entity.Product, error) {
	var product []entity.Product

	if err := r.db.Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindProductById(id int) (entity.Product, error) {
	var product entity.Product

	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) CreateProduct(product entity.Product) (entity.Product, error) {
	if err := r.db.Create(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

// func (r *repository) UpdateProduct(product entity.Product) (entity.Product, error) {
// 	if err := r.db.Save(&product).Error; err != nil {
// 		return product, err
// 	}
// 	return product, nil
// }

func (r *repository) UpdateProduct(id int, dataUpdate map[string]interface{}) (entity.Product, error) {
	var product entity.Product

	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	if err := r.db.Model(&product).Where("id = ?", id).Updates(dataUpdate).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) DeleteProduct(id int) (interface{}, error) {
	if err := r.db.Where("id = ?", id).Delete(&entity.Product{}).Error; err != nil {
		return "error", err
	}

	status := "product successfully deleted"

	return status, nil
}
