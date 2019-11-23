package rest

import (
	"encoding/json"

	"fasthttp-project/api/errors"
	"fasthttp-project/domain/model"
	"fasthttp-project/interface/api"
	"fasthttp-project/interface/service"
)

type User = model.User
type Users = model.Users
type UserOrders = model.UserOrders
type UserService = service.UserService

func GetUsersHandler(userService UserService) api.Handler {
	return func (ctx api.Context) (json.Marshaler, error) {
		offset, err := ctx.QueryOffsetArg()
		if err != nil {
			return nil, err
		}
		limit, err := ctx.QueryLimitArg()
		if err != nil {
			return nil, err
		}
		users, err := userService.GetUsers(offset, limit)
		if err != nil {
			return nil, err
		}
		return Users(users), nil
	}
}

func GetUserHandler(userService UserService) api.Handler {
	return func (ctx api.Context) (json.Marshaler, error) {
		userId, err := ctx.PathInt64Arg("user_id")
		if err != nil {
			return nil, errors.ErrUserId
		}
		user, err := userService.GetUser(userId)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func CreateUserHandler(userService UserService) api.Handler {
	return func (ctx api.Context) (json.Marshaler, error) {
		user := User{}
		err := user.UnmarshalJSON(ctx.Body())
		if err != nil {
			return nil, err
		}
		err = userService.CreateUser(&user)
		return nil, err
	}
}

func GetUserOrdersHandler(userService UserService) api.Handler {
	return func (ctx api.Context) (json.Marshaler, error) {
		offset, err := ctx.QueryOffsetArg()
		if err != nil {
			return nil, err
		}
		limit, err := ctx.QueryLimitArg()
		if err != nil {
			return nil, err
		}
		userOrders, err := userService.GetUserOrders(offset, limit)
		if err != nil {
			return nil, err
		}
		return UserOrders(userOrders), nil
	}
}
