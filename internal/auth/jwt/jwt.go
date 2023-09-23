package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"sso/internal/auth/types"
	"sso/internal/config"
	"sso/internal/user/types"
	"time"
)

func Generate(user types.User, allowedServices []string, duration time.Duration) (string, error) {
	secretKey := config.Get().Auth.JwtSecret

	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(duration).Unix(),
		IssuedAt:  time.Now().Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, auth.JWTClaims{
		StandardClaims:  standardClaims,
		AllowedServices: allowedServices,
		User:            user,
	})

	token, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func Parse(token string) (auth.JWTClaims, error) {
	secretKey := config.Get().Auth.JwtSecret

	jwtToken, err := jwt.ParseWithClaims(token, auth.JWTClaims{}, parse(secretKey))
	if err != nil {
		return auth.JWTClaims{}, err
	}

	claims, claimsOk := jwtToken.Claims.(auth.JWTClaims)
	if !claimsOk {
		// TODO: move to custom errors package
		return auth.JWTClaims{}, errors.New("invalid claims type")
	}

	if err = claims.Valid(); err != nil {
		return auth.JWTClaims{}, err
	}

	return claims, nil
}

func parse(secretKey string) jwt.Keyfunc {
	key := []byte(secretKey)

	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	}
}
