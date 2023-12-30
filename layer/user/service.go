package user

import (
	"errors"
	"fmt"
	"mini-commerce/entity"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	GetUserById(id int) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	SaveNewUser(user entity.UserInput) (entity.User, error)
	LoginUser(user entity.UserLogin) (entity.User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserById(id int) (entity.User, error) {
	user, err := s.repository.FindUserById(id)

	if err != nil {
		return entity.User{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user with id %d not found", id)
		return entity.User{}, errors.New(newError)
	}

	return user, nil
}

func (s *service) GetUserByEmail(email string) (entity.User, error) {
	user, err := s.repository.FindUserByEmail(email)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *service) SaveNewUser(user entity.UserInput) (entity.User, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	if err != nil {
		return entity.User{}, err
	}

	var NewUser = entity.User{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    string(genPassword),
		Role:        "user",
	}

	createUser, err := s.repository.Create(NewUser)
	if err != nil {
		return entity.User{}, err
	}
	return createUser, nil
}

func (s *service) LoginUser(input entity.UserLogin) (entity.User, error) {
	user, err := s.repository.FindUserByEmail(input.Email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user with id %v not found", user.ID)
		return entity.User{}, errors.New(newError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("invalid password")
	}

	return user, nil
}
