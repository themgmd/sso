package auth

import (
	"github.com/golang-jwt/jwt"
	"sso/internal/user/types"
)

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyTokenDTO struct {
	WithToken
	ServiceId string `json:"service_id"`
}

type LogoutDTO struct {
	WithToken
}

type WithToken struct {
	Token string `json:"token" uri:"token"`
}

type JWTClaims struct {
	jwt.StandardClaims
	AllowedServices []string
	User            types.User
}
