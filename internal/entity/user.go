package entity

import (
	"errors"
)

type (
	User struct {
		ID        int64  `json:"id" gorm:"column:id"`
		Email     string `json:"email"  gorm:"column:email"`
		Name      string `json:"name"  gorm:"column:name"`
		Username  string `json:"username"  gorm:"column:username"`
		Password  string `json:"password"  gorm:"column:password"`
		IsLogedIn bool   `json:"is_loged_in" gorm:"column:is_loged_in"`
		Seed      string `json:"seed" gorm:"column:seed"`
	}

	SignUpRequest struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	SignUpResponse struct {
		Message string `json:"message"`
	}

	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}

	LogoutRequest struct {
		Username string `json:"username"`
	}
)

func (s *SignUpRequest) Validate() error {
	if s.Email == "" {
		return errors.New("Email can not be empty")
	}

	if s.Username == "" {
		return errors.New("Username can not be empty")
	}

	if s.Password == "" {
		return errors.New("Password can not be empty")
	}

	return nil
}

func (l *LoginRequest) Validate() error {
	if l.Username == "" {
		return errors.New("Empty username")
	}

	if l.Password == "" {
		return errors.New("Empty password")
	}

	return nil
}
