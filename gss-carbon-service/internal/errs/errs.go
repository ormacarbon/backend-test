package errs

import (
	"errors"
	"fmt"
)

var (
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrInvalidReferralToken = errors.New("invalid referral token")
	ErrUserNotFound         = errors.New("user not found")
	ErrInternalServer       = errors.New("internal server error")
	ErrInvalidRequest       = errors.New("invalid request")
)

type AppError struct {
	Message    string
	StatusCode int
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(message string, statusCode int, err error) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}
