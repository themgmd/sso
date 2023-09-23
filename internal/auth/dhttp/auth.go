package dhttp

import (
	"context"
	"net/http"
	auth "sso/internal/auth/types"
	user "sso/internal/user/types"
	"sso/pkg/transport"

	"github.com/gorilla/mux"
)

type AuthService interface {
	Login(ctx context.Context, dto auth.LoginDTO) (string, error)
	Verify(ctx context.Context, dto auth.VerifyTokenDTO) (user.User, error)
	Logout(ctx context.Context, token string) error
}

type User struct {
	auth AuthService
}

func NewAuth(auth AuthService) *User {
	return &User{auth: auth}
}

func (u *User) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/auth/login", u.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", u.Verify).Methods(http.MethodPost)
	router.HandleFunc("/auth/logout", u.Logout).Methods(http.MethodPost)
}

func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	var dto auth.LoginDTO

	err := transport.ReadBody(r.Body, &dto)
	if err != nil {
		w.Write(transport.NewErrorResponse(err.Error()).Bytes())
		return
	}

	token, err := u.auth.Login(r.Context(), dto)
	if err != nil {
		w.Write(transport.NewErrorResponse(err.Error()).Bytes())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(transport.NewSuccessResponse(token).Bytes())
}

func (u *User) Verify(w http.ResponseWriter, r *http.Request) {
	var dto auth.VerifyTokenDTO

	err := transport.ReadBody(r.Body, &dto)
	if err != nil {
		w.Write(transport.NewErrorResponse(err.Error()).Bytes())
		return
	}

	customer, err := u.auth.Verify(r.Context(), dto)
	if err != nil {
		w.Write(transport.NewErrorResponse(err.Error()).Bytes())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(transport.NewSuccessResponse(customer).Bytes())
}

func (u *User) Logout(w http.ResponseWriter, r *http.Request) {
	var dto auth.LogoutDTO

	err := transport.ReadBody(r.Body, &dto)
	if err != nil {
		w.Write(transport.NewErrorResponse(err.Error()).Bytes())
		return
	}

	err = u.auth.Logout(r.Context(), dto.Token)
	if err != nil {
		w.Write(transport.NewErrorResponse(err.Error()).Bytes())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(transport.NewBaseResponse(true, "").Bytes())
}
