package user

import (
	"context"
	"sso/internal/user/types"
)

type Repo interface {
	GetByEmail(ctx context.Context, email string) (types.User, error)
}

type User struct {
	user Repo
}

func New(user Repo) *User {
	return &User{user: user}
}

func (u *User) GetByEmail(ctx context.Context, email string) (types.User, error) {
	return u.user.GetByEmail(ctx, email)
}
