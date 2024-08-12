package domain

import (
	"errors"
	"strings"
	"time"
)

type UserRole string

const (
	Admin   UserRole = "admin"
	Manager UserRole = "manager"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUserRole(role string) (UserRole, error) {
	switch strings.ToLower(role) {
	case string(Admin):
		return Admin, nil
	case string(Manager):
		return Manager, nil
	default:
		return "", errors.New("invalid user role")
	}
}
