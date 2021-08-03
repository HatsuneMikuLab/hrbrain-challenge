package models

import (
	"net/mail"
)

type User struct {
	User string `json:"user"`
	Email string `json:"email"`
}

func (u *User) Validate() []string {
	validationErrors := make([]string, 0, 2)
	if len(u.User) < 2 {
		validationErrors = append(validationErrors, "Name should be at least 2 characters long")
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		validationErrors = append(validationErrors, "Ivalid Email.")
	}
	return validationErrors
}