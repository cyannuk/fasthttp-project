package service

import (
	"io"

	"fasthttp-project/domain/model"
)

type OrderService interface {
	io.Closer
	GetOrder(id int64) (*model.Order, error)
	GetOrders(userId int64, offset int64, limit int64) ([]model.Order, error)
}
