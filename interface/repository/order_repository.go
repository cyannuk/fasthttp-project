package repository

import (
	"fasthttp-project/domain/model"
)

type OrderRepository interface {
	GetOrder(id int64) (*model.Order, error)
	GetOrders(userId int64, offset int64, limit int64) ([]model.Order, error)
}
