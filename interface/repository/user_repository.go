package repository

import (
	"io"

	"fasthttp-project/domain/model"
)

type UserRepository interface {
	io.Closer
	GetUser(id int64) (*model.User, error)
	GetUsers(offset int64, limit int64) ([]model.User, error)
	GetUserOrders(offset int64, limit int64) ([]model.UserOrder, error)
	CreateUser(user *model.User) error
}
