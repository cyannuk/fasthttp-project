package repository

import (
	"fasthttp-project/domain/model"
)

type User = model.User
type UserOrder = model.UserOrder

type UserRepository interface {
	GetUser(id int64) (*User, error)
	GetUsers(offset int, limit int) (users []User, err error)
	GetUserOrders(offset int, limit int) (userOrders []UserOrder, err error)
	CreateUser(user *User) error
}
