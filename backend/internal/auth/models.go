package auth

import (
	"time"
)

type Role string

const (
	Admin    Role = "admin"
	Customer Role = "customer"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Role           Role
	CreatedAt      time.Time
}

type UserSignupForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
