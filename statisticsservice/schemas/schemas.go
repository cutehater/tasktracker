package schemas

type EventType int

const (
	View EventType = iota
	Like
)

type Event struct {
	TaskID        int64     `json:"task_id"`
	UserID        int64     `json:"user_id"`
	OwnerUsername string    `json:"owner_username"`
	EventType     EventType `json:"event_type"`
}
