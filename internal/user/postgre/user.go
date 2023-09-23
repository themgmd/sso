package postgre

import (
	"context"
	"sso/internal/user/types"
	"sso/pkg/connectors/postgre"
)

type User struct {
	db *postgre.DB
}

func NewUser(db *postgre.DB) *User {
	return &User{db: db}
}

func (u *User) GetByEmail(ctx context.Context, email string) (types.User, error) {
	var user types.User

	_, err := u.db.ExecContext(ctx, queryGetUserByEmail, email)
	if err != nil {
		return user, err
	}

	return user, nil
}