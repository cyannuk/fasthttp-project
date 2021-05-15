package api

import (
	"fasthttp-project/api/errors"
	"fasthttp-project/domain/model"
	"github.com/goccy/go-json"
	"github.com/valyala/fasthttp"
)

func (application *Application) getUsers(ctx *fasthttp.RequestCtx) {
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
	users, err := application.userService.GetUsers(offset, limit)
	if err != nil {
		sendError(err, ctx)
		return
	}
	sendData(users, ctx)
}

func (application *Application) getUser(ctx *fasthttp.RequestCtx) {
	userId, err := getPathParameter(ctx, "user_id")
	if err != nil || userId <= 0 {
		sendError(errors.ErrUserId, ctx)
		return
	}
	user, err := application.userService.GetUser(userId)
	if err != nil {
		sendError(err, ctx)
		return
	}
	sendData(user, ctx)
}

func (application *Application) createUser(ctx *fasthttp.RequestCtx) {
	var user = model.User{}
	err := json.Unmarshal(ctx.Request.Body(), &user)
	if err != nil {
		sendError(err, ctx)
		return
	}
	err = application.userService.CreateUser(&user)
	if err != nil {
		sendError(err, ctx)
		return
	}
	sendData(user.ID, ctx)
}

func (application *Application) getUserOrders(ctx *fasthttp.RequestCtx) {
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
	userOrders, err := application.userService.GetUserOrders(offset, limit)
	if err != nil {
		sendError(err, ctx)
		return
	}
	sendData(userOrders, ctx)
}
