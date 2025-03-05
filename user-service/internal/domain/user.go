package domain

import (
	"fmt"
	"strings"
)

type User struct {
	id       int
	email    string
	password string
}

func NewUser(id int, email, password string) (*User, error) {
	user := &User{id: id, email: email, password: password}

	// if err := user.ValidateUser(); err != nil {
	// 	return nil, err
	// }

	return user, nil
}

func CreateUser(email, password string) (*User, error) {
	user := &User{id: 0, email: email, password: password}

	if err := user.ValidateUser(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) ID() int {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) ChangeEmail(email string) error {
	if err := ValidateEmail(email); err != nil {
		return err
	}

	u.email = email

	return nil
}

func (u *User) ChangePassword(password string) error {
	if err := ValidatePassword(password); err != nil {
		return err
	}

	u.password = password

	return nil
}

func ValidatePassword(password string) error {
	if len := len(password); len < 8 || len > 32 {
		return fmt.Errorf("%w: len < 8 or len > 32", ErrPasswordInvalid)
	}
	return nil
}

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf("%w: @ not contains", ErrEmailInvalid)
	}
	return nil
}

func (u *User) ValidateUser() error {
	var err error
	if err = ValidateEmail(u.email); err != nil {
		return err
	}

	if err = ValidatePassword(u.password); err != nil {
		return err
	}
	return nil
}
