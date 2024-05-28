package schemas

import (
	"gorm.io/gorm"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type EventType int

const (
	View EventType = iota
	Like
)

type Event struct {
	TaskID    int64     `json:"task_id"`
	UserID    int64     `json:"user_id"`
	EventType EventType `json:"event_type"`
}

type UserData struct {
	gorm.Model
	Login        string `json:"login,omitempty"`
	PasswordHash string `json:"password,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	BirthDate    string `json:"birthDate,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
}
