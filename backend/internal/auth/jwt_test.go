package auth

import (
	"testing"
	"time"
)

func TestGenerateAccessToken(t *testing.T) {
	user := User{
		ID: 1,
		Name: "Bob",
		Email: "bob@test.com",
		HashedPassword: []byte("someHashedPassword"),
		Role: Customer,
		CreatedAt: time.Now(),
	}

	_, err := generateAccessToken(&user)
	if err != nil {
		t.Fatal(err)
	}
}
