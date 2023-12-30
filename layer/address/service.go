package address

import (
	"errors"
	"fmt"
	"mini-commerce/entity"
)

type Service interface {
	GetAddressByUserId(userId int) ([]entity.Address, error)
	GetAddressById(addressId int) (entity.Address, error)
	SaveNewAddress(userId int, input entity.AddressInput) (entity.Address, error)
	UpdateAddress(addressId int, dataUpdate entity.AddressUpdate) (entity.Address, error)
	DeleteAddress(addressId int) (interface{}, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAddressByUserId(userId int) ([]entity.Address, error) {
	addresses, err := s.repository.FindAddressByUserId(userId)
	if err != nil {
		return []entity.Address{}, err
	}
	return addresses, nil
}

func (s *service) GetAddressById(addressId int) (entity.Address, error) {
	address, err := s.repository.FindAddressById(addressId)
	if err != nil {
		return entity.Address{}, err
	}

	if address.ID == 0 {
		newError := fmt.Sprintf("address with id %d not found", addressId)
		return entity.Address{}, errors.New(newError)
	}
	return address, nil
}

func (s *service) SaveNewAddress(userId int, input entity.AddressInput) (entity.Address, error) {
	address := entity.Address{
		Receiver:      input.Receiver,
		PhoneReceiver: input.PhoneReceiver,
		AddressDetail: input.AddressDetail,
		Province:      input.Province,
		City:          input.City,
		UserID:        userId,
	}
	newAddress, err := s.repository.CreateAddress(address)
	if err != nil {
		return entity.Address{}, err
	}
	return newAddress, nil
}

func (s *service) UpdateAddress(addressId int, inputUpdate entity.AddressUpdate) (entity.Address, error) {
	dataUpdate := map[string]interface{}{}

	address, err := s.repository.FindAddressById(addressId)
	if err != nil {
		return entity.Address{}, err
	}
	if address.ID == 0 {
		newError := fmt.Sprintf("address with id %d not found", addressId)
		return entity.Address{}, errors.New(newError)
	}

	if inputUpdate.Receiver != "" || len(inputUpdate.Receiver) != 0 {
		dataUpdate["receiver"] = inputUpdate.Receiver
	}

	if inputUpdate.PhoneReceiver != "" || len(inputUpdate.PhoneReceiver) != 0 {
		dataUpdate["phone_receiver"] = inputUpdate.PhoneReceiver
	}

	if inputUpdate.AddressDetail != "" || len(inputUpdate.AddressDetail) != 0 {
		dataUpdate["address_detail"] = inputUpdate.AddressDetail
	}

	if inputUpdate.Province != "" || len(inputUpdate.Province) != 0 {
		dataUpdate["province"] = inputUpdate.Province
	}

	if inputUpdate.City != "" || len(inputUpdate.City) != 0 {
		dataUpdate["city"] = inputUpdate.City
	}

	addressUpdated, err := s.repository.UpdateByAddressId(addressId, dataUpdate)
	if err != nil {
		return entity.Address{}, err
	}
	return addressUpdated, nil
}

func (s *service) DeleteAddress(addressId int) (interface{}, error) {
	address, err := s.repository.FindAddressById(addressId)
	if err != nil {
		return nil, err
	}
	if address.ID == 0 {
		newError := fmt.Sprintf("address with id %d not found", addressId)
		return nil, errors.New(newError)
	}

	statusDelete, _ := s.repository.DeleteByAddressId(addressId)
	if statusDelete == "error" {
		return nil, errors.New("error delete in internal server")
	}

	message := fmt.Sprintf("address with id %d has been deleted", addressId)
	return message, nil
}
