package auth

import (
	"context"
	"strconv"
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo *repository
}

func NewService(r *repository) *service {
	return &service{repo: r}
}

// Create new user using form credentials. Return error if bad credentials.
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

// Authenticate login credentials. Return AuthResult struct if successful. Else, return error
func (s *service) Authenticate(ctx context.Context, form UserLoginForm) (*AuthResult, error) {
	user, err := s.repo.Fetch(ctx, form.Email)
	if err != nil {
		return nil, err
	}

	// validate hashed password
	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(form.Password))
	if err != nil {
		return nil, ErrBadCredentials
	}

	userInfo := UserInfo{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
	}

	// generate JWT
	accessToken, err := generateAccessToken(user)
	if err != nil {
		return nil, err
	}

	authResult := AuthResult{
		AccessToken: accessToken,
		RefreshToken: "",
		UserInfo: userInfo,
	}

	return &authResult, nil
}

// Generates JWT access token
func generateAccessToken(user *User) (string, error) {
	type CustomClaims struct {
		Role Role `json:"role"`
		jwt.RegisteredClaims
	}

	// access token with 1h expiry date
	claims := CustomClaims{
		user.Role,
		jwt.RegisteredClaims{
			Issuer: "go-comm-api",
			Subject: strconv.Itoa(user.ID),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := []byte(os.Getenv("ACCESS_SECRET_KEY"))

	s, err := t.SignedString(secret)
	if err != nil {
		return "", err
	}
	
	return s, nil
}
