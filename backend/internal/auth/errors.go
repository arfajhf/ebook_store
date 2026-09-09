package auth

import "errors"

var (
	ErrEmailAlreadyUsed = errors.New("email already used")
	ErrInvalidName      = errors.New("invalid name")
)
