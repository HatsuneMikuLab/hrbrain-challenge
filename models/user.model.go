package models

import (
	"net/mail"
)

type User struct {
	ID string `json:"user"`
	Email string `json:"email"`
}

type DisrespectedUser struct {
	ID string `json:"user"`
	Evaluation string `json:"evaluation"`
}

func (u *User) Validate() []string {
	validationErrors := make([]string, 0, 2)
	if len(u.ID) < 2 {
		validationErrors = append(validationErrors, "ID (user) should be at least 2 characters long")
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		validationErrors = append(validationErrors, "Ivalid Email.")
	}
	return validationErrors
}