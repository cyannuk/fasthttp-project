package errors

import (
	"errors"
)

var (
	ErrOffset = errors.New("invalid offset value")
	ErrLimit = errors.New("invalid limit value")
	ErrUserId = errors.New("invalid user_id value")
	ErrOrderId = errors.New("invalid order_id value")
	ErrForbidden = errors.New("forbidden resource")
)
