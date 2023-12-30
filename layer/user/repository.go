package user

import (
	"mini-commerce/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindUserByEmail(email string) (entity.User, error)
	Create(user entity.User) (entity.User, error)
	FindUserById(id int) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUserByEmail(email string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("email = ?", email).Preload("Products").Preload("Address").First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Create(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindUserById(id int) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", id).Preload("Products").Preload("Address").First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
