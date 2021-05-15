package api

import (
	"fasthttp-project/api/errors"
	"github.com/valyala/fasthttp"
)

func (application *Application) getOrders(ctx *fasthttp.RequestCtx) {
	userId, err := getPathParameter(ctx, "user_id")
	if err != nil || userId <= 0 {
		sendError(errors.ErrUserId, ctx)
		return
	}
	offset, err := getQueryOffset(ctx)
	if err != nil {
		sendError(err, ctx)
		return
	}
	limit, err := getQueryLimit(ctx)
	if err != nil {
		sendError(err, ctx)
		return
	}
	orders, err := application.orderService.GetOrders(userId, offset, limit)
	if err != nil {
		sendError(err, ctx)
		return
	}
	sendData(orders, ctx)
}

func (application *Application) getUserOrder(ctx *fasthttp.RequestCtx) {
	userId, err := getPathParameter(ctx, "user_id")
	if err != nil || userId <= 0 {
		sendError(errors.ErrUserId, ctx)
		return
	}
	orderId, err := getPathParameter(ctx, "order_id")
	if err != nil || orderId <= 0 {
		sendError(errors.ErrOrderId, ctx)
		return
	}
	order, err := application.orderService.GetOrder(orderId)
	if err != nil {
		sendError(err, ctx)
		return
	}
	if order.UserID != userId {
		sendError(errors.ErrForbidden, ctx)
		return
	}
	sendData(order, ctx)
}
