package service

import (
	"fasthttp-project/interface/repository"
	"fasthttp-project/interface/service"
)

type Order = service.Order

type OrderService struct {
	repository.OrderRepository
}

func (o OrderService) GetOrder(id int64) (*Order, error) {
	return o.OrderRepository.GetOrder(id)
}

func (o OrderService) GetOrders(userId int64, offset int, limit int) (orders []Order, err error) {
	return o.OrderRepository.GetOrders(userId, offset, limit)
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return OrderService{repo}
}
