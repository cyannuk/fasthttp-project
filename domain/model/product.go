package model

import (
	"time"
)

//go:generate easyjson -snake_case -omit_empty
//go:generate reform

//reform:products
//easyjson:json
type Product struct {
	ID        int64     `reform:"id,pk"`
	CreatedAt time.Time `reform:"created_at"`
	Category  string    `reform:"category"`
	Ean       string    `reform:"ean"`
	Price     float64   `reform:"price"`
	Quantity  int32     `reform:"quantity"`
	Rating    float64   `reform:"rating"`
	Title     string    `reform:"title"`
	Vendor    string    `reform:"vendor"`
}

//easyjson:json
type Products []Product
