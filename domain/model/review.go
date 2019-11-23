package model

import (
	"time"
)

//go:generate easyjson -snake_case -omit_empty
//go:generate reform

//reform:reviews
//easyjson:json
type Review struct {
	ID        int64     `reform:"id,pk"`
	CreatedAt time.Time `reform:"created_at"`
	Reviewer  string    `reform:"reviewer"`
	ProductID int64     `reform:"product_id"`
	Rating    int32     `reform:"rating"`
	Body      string    `reform:"body"`
}

//easyjson:json
type Reviews []Review
