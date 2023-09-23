package dhttp

import (
	"context"
	"sso/internal/user/types"
)

type UserService interface {
	GetByEmail(ctx context.Context, email string) (types.User, error)
}

type User struct {
	user UserService
}

func NewUser(user UserService) *User {
	return &User{user: user}
}
