package domain

type FriendStatus string

const (
	StatusPending  FriendStatus = "pending"
	StatusAccepted FriendStatus = "accepted"
)

func (s FriendStatus) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusAccepted:
		return "accepted"
	default:
		return "unknown"
	}
}
