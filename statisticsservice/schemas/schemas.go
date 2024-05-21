package schemas

type EventType int

const (
	View EventType = iota
	Like
)

type Event struct {
	TaskId    int64     `json:"task_id"`
	UserId    int64     `json:"user_id"`
	EventType EventType `json:"event_type"`
}
