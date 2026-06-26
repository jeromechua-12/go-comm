package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo *repository
}

func NewService(r *repository) *service {
	return &service{repo: r}
}

// Create new user with hashed password in db
func (s *service) Signup(ctx context.Context, form UserSignupForm) error {
	// encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), 12)
	if err != nil {
		return err
	}

	// insert into db
	err = s.repo.Insert(ctx, form.Name, form.Email, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
