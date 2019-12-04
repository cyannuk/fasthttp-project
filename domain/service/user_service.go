package service

import (
	"fasthttp-project/domain/model"
	"fasthttp-project/interface/repository"
)

type User = model.User
type UserOrder = model.UserOrder

type UserService struct {
	repository.UserRepository
}

func (u UserService) GetUser(id int64) (*User, error) {
	return u.UserRepository.GetUser(id)
}

func (u UserService) GetUsers(offset int, limit int) (users []User, err error) {
	return u.UserRepository.GetUsers(offset, limit)
}

func (u UserService) GetUserOrders(offset int, limit int) (userOrders []UserOrder, err error) {
	return u.UserRepository.GetUserOrders(offset, limit)
}

func (u UserService) CreateUser(user *User) error {
	return u.UserRepository.CreateUser(user)
}

func NewUserService(repo repository.UserRepository) UserService {
	return UserService{repo}
}
