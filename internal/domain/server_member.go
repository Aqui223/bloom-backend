package domain

import "time"

type ServerMember struct {
	ID       int       `json:"id" db:"id"`
	ServerID int       `json:"server_id" db:"server_id"`
	MemberID int       `json:"member_id" db:"member_id"`
	JoinedAt time.Time `json:"joined_at" db:"joined_at"`
}
