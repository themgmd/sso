package auth

import (
	"context"
	"errors"
	"slices"
	"sso/internal/auth/jwt"
	auth "sso/internal/auth/types"
	"sso/internal/user/types"
	"time"
)

type UserService interface {
	GetByEmail(ctx context.Context, email string) (types.User, error)
}

type Auth struct {
	user UserService
}

func New(user UserService) *Auth {
	return &Auth{user: user}
}

func (a *Auth) Login(ctx context.Context, dto auth.LoginDTO) (string, error) {
	user, err := a.user.GetByEmail(ctx, dto.Email)
	if err != nil {
		return "", err
	}

	err = user.ComparePassword(dto.Password)
	if err != nil {
		// TODO: make errors package
		return "", errors.New("invalid credentials")
	}

	return jwt.Generate(user, []string{}, time.Minute*30)
}

func (a *Auth) Verify(_ context.Context, dto auth.VerifyTokenDTO) (types.User, error) {
	claims, err := jwt.Parse(dto.Token)
	if err != nil {
		return types.User{}, err
	}

	contains := slices.Contains(claims.AllowedServices, dto.ServiceId)
	if !contains {
		return types.User{}, errors.New("forbidden")
	}

	return claims.User, nil
}

func (a *Auth) Logout(ctx context.Context, token string) error {
	return nil
}
