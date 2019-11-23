package rest

import (
	"encoding/json"

	"fasthttp-project/api/errors"
	"fasthttp-project/domain/model"
	"fasthttp-project/interface/api"
	"fasthttp-project/interface/service"
)

type Order = model.Order
type Orders = model.Orders
type OrderService = service.OrderService

func GetOrdersHandler(orderService OrderService) api.Handler {
	return func (ctx api.Context) (json.Marshaler, error) {
		userId, err := ctx.PathInt64Arg("user_id")
		if err != nil {
			return nil, errors.ErrUserId
		}
		offset, err := ctx.QueryOffsetArg()
		if err != nil {
			return nil, err
		}
		limit, err := ctx.QueryLimitArg()
		if err != nil {
			return nil, err
		}
		orders, err := orderService.GetOrders(userId, offset, limit)
		if err != nil {
			return nil, err
		}
		return Orders(orders), nil
	}
}

func GetOrderHandler(orderService OrderService) api.Handler {
	return func (ctx api.Context) (json.Marshaler, error) {
		userId, err := ctx.PathInt64Arg("user_id")
		if err != nil {
			return nil, errors.ErrUserId
		}
		orderId, err := ctx.PathInt64Arg("order_id")
		if err != nil {
			return nil, errors.ErrOrderId
		}
		order, err := orderService.GetOrder(orderId)
		if err != nil {
			return nil, err
		}
		if order.UserID != userId {
			return nil, errors.ErrForbidden
		}
		return order, nil
	}
}
