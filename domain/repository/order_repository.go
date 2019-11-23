package repository

import (
	"gopkg.in/reform.v1"

	"fasthttp-project/domain/model"
)

type Order = model.Order

type orderRepository struct {
	dataSource
}

func (orderRepo orderRepository) GetOrder(id int64) (*Order, error) {
	order, err := orderRepo.FindByPrimaryKeyFrom(model.OrderTable, id)
	if err != nil {
		return nil, err
	}
	return order.(*Order), nil
}

func (orderRepo orderRepository) GetOrders(userId int64, offset int, limit int) (orders []Order, err error) {
	rows, err := orderRepo.SelectRows(model.OrderTable, "WHERE user_id = $1 ORDER BY id OFFSET $2 LIMIT $3",
		userId, offset, limit)
	if err != nil {
		return
	}
	defer func() {
		e := rows.Close()
		if e != nil && err == nil {
			err = e
		}
	}()
	orders = make([]Order, 0, limit)
	for {
		var order Order
		if err = orderRepo.NextRow(&order, rows); err != nil {
			break
		}
		orders = append(orders, order)
	}
	if err == reform.ErrNoRows {
		err = nil
	}
	return
}

func NewOrderRepository(db dataSource) orderRepository {
	return orderRepository{db}
}
