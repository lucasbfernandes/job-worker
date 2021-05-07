package user

import (
	"github.com/google/uuid"

	"fmt"
)

const (
	AdminRole = "ADMIN"
	UserRole  = "USER"
)

type User struct {
	ID       string
	Username string
	APIToken string
	Role     string
}

func NewUser(username string, apiToken string, role string) (*User, error) {
	if role != AdminRole && role != UserRole {
		return nil, fmt.Errorf("invalid role %s", role)
	}

	return &User{
		ID:       uuid.New().String(),
		Username: username,
		APIToken: apiToken,
		Role:     role,
	}, nil
}
