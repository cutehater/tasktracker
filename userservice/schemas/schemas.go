package schemas

import (
	"gorm.io/gorm"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Like struct {
	TaskId int `json:"task_id"`
	UserId int `json:"user_id"`
}

type View struct {
	TaskId int `json:"task_id"`
	UserId int `json:"user_id"`
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
