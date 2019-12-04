package repository

import (
	"gopkg.in/reform.v1"

	"fasthttp-project/domain/model"
)

type User = model.User
type UserOrder = model.UserOrder

type UserRepository struct {
	DataSource
}

func (userRepo UserRepository) GetUser(id int64) (*User, error) {
	user, err := userRepo.FindByPrimaryKeyFrom(model.UserTable, id)
	if err != nil {
		return nil, err
	}
	return user.(*User), nil
}

func (userRepo UserRepository) GetUsers(offset int, limit int) (users []model.User, err error) {
	rows, err := userRepo.SelectRows(model.UserTable, "ORDER BY id OFFSET $1 LIMIT $2", offset, limit)
	if err != nil {
		return
	}
	defer func() {
		e := rows.Close()
		if e != nil && err == nil {
			err = e
		}
	}()
	users = make([]User, 0, limit)
	for {
		var user User
		if err = userRepo.NextRow(&user, rows); err != nil {
			break
		}
		users = append(users, user)
	}
	if err == reform.ErrNoRows {
		err = nil
	}
	return
}

func (userRepo UserRepository) GetUserOrders(offset int, limit int) (userOrders []UserOrder, err error) {
	rows, err := userRepo.Query(
		"SELECT u.name, u.city, u.state, o.product_id, o.quantity, o.total " +
		"FROM users u " +
		"INNER JOIN orders o ON o.user_id = u.id " +
		"ORDER BY u.id, o.id " +
		"OFFSET $1 LIMIT $2", offset, limit)
	if err != nil {
		return
	}
	defer func() {
		e := rows.Close()
		if e != nil && err == nil {
			err = e
		}
	}()
	userOrders = make([]UserOrder, 0, limit)
	for {
		var userOrder UserOrder
		if err = userRepo.NextRow(&userOrder, rows); err != nil {
			break
		}
		userOrders = append(userOrders, userOrder)
	}
	if err == reform.ErrNoRows {
		err = nil
	}
	return
}

func (userRepo UserRepository) CreateUser(user *User) error {
	return userRepo.Insert(user)
}

func NewUserRepository(ds DataSource) UserRepository {
	return UserRepository{ds}
}
