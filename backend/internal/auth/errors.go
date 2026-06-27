package auth

import "errors"

var (
	ErrDuplicateEmail = errors.New("user: duplicate email")
	ErrBadCredentials = errors.New("user: bad credential")
	ErrCookieTooLong  = errors.New("cookie: cookie value too long")
)
