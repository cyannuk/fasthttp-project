package rest

import (
	"encoding/json"

	"fasthttp-project/api"
	"fasthttp-project/api/errors"
	"fasthttp-project/domain/model"
	"fasthttp-project/interface/service"
)

type User = model.User
type Users = model.Users
type UserOrders = model.UserOrders
type UserService = service.UserService

func GetUsersHandler(userService UserService) api.Handler {
	return func (ctx api.Context) (result json.Marshaler, err error) {
		offset, err := ctx.QueryOffsetArg()
		if err != nil {
			return
		}
		limit, err := ctx.QueryLimitArg()
		if err != nil {
			return
		}
		users, err := userService.GetUsers(offset, limit)
		if err == nil {
			result = Users(users)
		}
		return
	}
}

func GetUserHandler(userService UserService) api.Handler {
	return func (ctx api.Context) (result json.Marshaler, err error) {
		userId, err := ctx.PathInt64Arg("user_id")
		if err != nil {
			err = errors.ErrUserId
			return
		}
		user, err := userService.GetUser(userId)
		if err == nil {
			result = user
		}
		return
	}
}

func CreateUserHandler(userService UserService) api.Handler {
	return func (ctx api.Context) (_ json.Marshaler, err error) {
		user := User{}
		err = user.UnmarshalJSON(ctx.Request.Body())
		if err == nil {
			err = userService.CreateUser(&user)
		}
		return
	}
}

func GetUserOrdersHandler(userService UserService) api.Handler {
	return func(ctx api.Context) (result json.Marshaler, err error) {
		offset, err := ctx.QueryOffsetArg()
		if err != nil {
			return
		}
		limit, err := ctx.QueryLimitArg()
		if err != nil {
			return
		}
		userOrders, err := userService.GetUserOrders(offset, limit)
		if err == nil {
			result = UserOrders(userOrders)
		}
		return
	}
}
