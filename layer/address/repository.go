package address

import (
	"mini-commerce/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAddressByUserId(userId int) ([]entity.Address, error)
	FindAddressById(addressId int) (entity.Address, error)
	CreateAddress(address entity.Address) (entity.Address, error)
	UpdateByAddressId(addressId int, dataUpdate map[string]interface{}) (entity.Address, error)
	DeleteByAddressId(addressId int) (interface{}, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAddressByUserId(userId int) ([]entity.Address, error) {
	var addresses []entity.Address
	if err := r.db.Where("user_id = ?", userId).Find(&addresses).Error; err != nil {
		return addresses, err
	}
	return addresses, nil
}

func (r *repository) FindAddressById(addressId int) (entity.Address, error) {
	var address entity.Address
	if err := r.db.Where("id = ?", addressId).First(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}

func (r *repository) CreateAddress(address entity.Address) (entity.Address, error) {
	if err := r.db.Create(&address).Error; err != nil {
		return address, err
	}
	return address, nil
}

func (r *repository) UpdateByAddressId(addressId int, dataUpdate map[string]interface{}) (entity.Address, error) {
	var address entity.Address
	if err := r.db.Where("id = ?", addressId).First(&address).Error; err != nil {
		return address, err
	}
	if err := r.db.Model(&address).Where("id = ?", addressId).Updates(dataUpdate).Error; err != nil {
		return address, err
	}
	return address, nil
}

func (r *repository) DeleteByAddressId(addressId int) (interface{}, error) {
	if err := r.db.Where("id = ?", addressId).Delete(&entity.Address{}).Error; err != nil {
		return "error", err
	}
	status := "address successfully deleted"
	return status, nil
}
