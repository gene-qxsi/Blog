package domain

import "errors"

var (
	ErrCreatedUser     = errors.New("invalid user")
	ErrEmailInvalid    = errors.New("email invalid")
	ErrPasswordInvalid = errors.New("password invalid")
)
