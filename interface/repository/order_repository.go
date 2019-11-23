package repository

import (
	"fasthttp-project/domain/model"
)

type Order = model.Order

type OrderRepository interface {
	GetOrder(id int64) (*Order, error)
	GetOrders(userId int64, offset int, limit int) (orders []Order, err error)
}
