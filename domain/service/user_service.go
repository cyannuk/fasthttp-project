package service

import (
	"fasthttp-project/domain/model"
	"fasthttp-project/interface/repository"
)

type User = model.User
type UserOrder = model.UserOrder

type userService struct {
	repository.UserRepository
}

func (u userService) GetUser(id int64) (*User, error) {
	return u.UserRepository.GetUser(id)
}

func (u userService) GetUsers(offset int, limit int) (users []User, err error) {
	return u.UserRepository.GetUsers(offset, limit)
}

func (u userService) GetUserOrders(offset int, limit int) (userOrders []UserOrder, err error) {
	return u.UserRepository.GetUserOrders(offset, limit)
}

func (u userService) CreateUser(user *User) error {
	return u.UserRepository.CreateUser(user)
}

func NewUserService(repo repository.UserRepository) userService {
	return userService{repo}
}
