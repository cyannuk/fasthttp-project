package repository

import (
	"fasthttp-project/domain/model"
)

type UserRepository interface {
	GetUser(id int64) (*model.User, error)
	GetUsers(offset int64, limit int64) ([]model.User, error)
	GetUserOrders(offset int64, limit int64) ([]model.UserOrder, error)
	CreateUser(user *model.User) error
}
