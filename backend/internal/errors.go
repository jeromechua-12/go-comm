package internal

// internal error types
type ErrorType string

const (
	ErrInternal        ErrorType = "internal_server_error"
	ErrBadRequest      ErrorType = "bad_request"
	ErrValidation      ErrorType = "validation_error"
	ErrBadCredentials  ErrorType = "bad_credentials"
	ErrUnauthorized    ErrorType = "unauthorized"
	ErrAccessExpired   ErrorType = "access_token_expired"
)
