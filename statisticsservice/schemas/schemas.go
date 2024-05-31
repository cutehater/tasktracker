package schemas

type EventType int

const (
	View EventType = iota
	Like
)

type Event struct {
	TaskID    int64     `json:"task_id"`
	Username  string    `json:"username"`
	EventType EventType `json:"event_type"`
}
