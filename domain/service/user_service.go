package service

import (
	"fasthttp-project/domain/model"
	"fasthttp-project/interface/repository"
)

type userService struct {
	repository.UserRepository
}

func (service userService) GetUser(id int64) (*model.User, error) {
	return service.UserRepository.GetUser(id)
}

func (service userService) GetUsers(offset int64, limit int64) ([]model.User, error) {
	return service.UserRepository.GetUsers(offset, limit)
}

func (service userService) GetUserOrders(offset int64, limit int64) ([]model.UserOrder, error) {
	return service.UserRepository.GetUserOrders(offset, limit)
}

func (service userService) CreateUser(user *model.User) error {
	return service.UserRepository.CreateUser(user)
}

func NewUserService(repository repository.UserRepository) userService {
	return userService{repository}
}
