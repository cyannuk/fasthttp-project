package service

import (
	"fasthttp-project/domain/model"
	"fasthttp-project/interface/repository"
	"fasthttp-project/interface/service"
)

type orderService struct {
	repository.OrderRepository
}

func (service orderService) GetOrder(id int64) (*model.Order, error) {
	return service.OrderRepository.GetOrder(id)
}

func (service orderService) GetOrders(userId int64, offset int64, limit int64) ([]model.Order, error) {
	return service.OrderRepository.GetOrders(userId, offset, limit)
}

func NewOrderService(repository repository.OrderRepository) service.OrderService {
	return orderService{repository}
}
