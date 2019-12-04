package rest

import (
	"encoding/json"

	"fasthttp-project/api"
	"fasthttp-project/api/errors"
	"fasthttp-project/domain/model"
	"fasthttp-project/interface/service"
)

type Order = model.Order
type Orders = model.Orders
type OrderService = service.OrderService

func GetOrdersHandler(orderService OrderService) api.Handler {
	return func (ctx api.Context) (result json.Marshaler, err error) {
		userId, err := ctx.PathInt64Arg("user_id")
		if err != nil {
			err = errors.ErrUserId
			return
		}
		offset, err := ctx.QueryOffsetArg()
		if err != nil {
			return
		}
		limit, err := ctx.QueryLimitArg()
		if err != nil {
			return
		}
		orders, err := orderService.GetOrders(userId, offset, limit)
		if err == nil {
			result = Orders(orders)
		}
		return
	}
}

func GetOrderHandler(orderService OrderService) api.Handler {
	return func (ctx api.Context) (result json.Marshaler, err error) {
		userId, err := ctx.PathInt64Arg("user_id")
		if err != nil {
			err = errors.ErrUserId
			return
		}
		orderId, err := ctx.PathInt64Arg("order_id")
		if err != nil {
			err = errors.ErrOrderId
			return
		}
		order, err := orderService.GetOrder(orderId)
		if err != nil {
			return
		}
		if order.UserID != userId {
			err = errors.ErrForbidden
		} else {
			result = order
		}
		return
	}
}
