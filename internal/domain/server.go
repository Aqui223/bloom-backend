package domain

import "time"

type Server struct {
	ID          int       `json:"id" db:"id"`
	OwnerID     int       `json:"owner_id" db:"owner_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
}
