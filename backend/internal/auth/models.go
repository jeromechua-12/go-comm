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

type UserLoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResult struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	UserInfo     UserInfo `json:"user_info"`
}

type UserInfo struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}
