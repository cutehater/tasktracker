package schemas

import (
	"gorm.io/gorm"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserData struct {
	gorm.Model
	Login        string `json:"login,omitempty"`
	PasswordHash string `json:"password, omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
	BirthDate    string `json:"birthDate,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
}
