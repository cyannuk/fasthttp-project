package model

import (
	"time"
)

//go:generate reform

//reform:orders
type Order struct {
	ID        int64     `reform:"id,pk"`
	CreatedAt time.Time `reform:"created_at"`
	UserID    int64     `reform:"user_id"`
	ProductID int64     `reform:"product_id"`
	Discount  *float64  `reform:"discount"`
	Quantity  int32     `reform:"quantity"`
	Subtotal  float64   `reform:"subtotal"`
	Tax       float64   `reform:"tax"`
	Total     float64   `reform:"total"`
}
