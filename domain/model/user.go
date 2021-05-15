package model

import (
	"time"
)

//go:generate reform

//reform:users
type User struct {
	ID        int64     `reform:"id,pk" json:",omit-dec"`
	CreatedAt time.Time `reform:"created_at"`
	Name      string    `reform:"name"`
	Email     string    `reform:"email"`
	Address   string    `reform:"address"`
	City      string    `reform:"city"`
	State     string    `reform:"state"`
	Zip       string    `reform:"zip"`
	BirthDate string    `reform:"birth_date"`
	Latitude  float64   `reform:"latitude"`
	Longitude float64   `reform:"longitude"`
	Password  string    `reform:"password" json:",omit-enc"`
	Source    string    `reform:"source"`
}
