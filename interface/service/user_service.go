package service

import (
	"fasthttp-project/domain/model"
)

type User = model.User
type UserOrder = model.UserOrder

type UserService interface {
	GetUser(id int64) (*User, error)
	GetUsers(offset int, limit int) (users []User, err error)
	GetUserOrders(offset int, limit int) (userOrders []UserOrder, err error)
	CreateUser(user *User) error
}
