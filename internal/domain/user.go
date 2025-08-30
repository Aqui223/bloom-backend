package domain

import "time"

type User struct {
	ID       int       `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"password"`
	Date     time.Time `db:"date" json:"date"`
}
