package models

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"-"`
	IsAdmin      bool   `json:"is_admin"`
}
