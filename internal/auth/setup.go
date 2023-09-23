package auth

import (
	"github.com/gorilla/mux"
	"sso/internal/auth/dhttp"
	"sso/internal/user"
)

func Setup(user *user.User, router *mux.Router) {
	auth := New(user)
	handlers := dhttp.NewAuth(auth)
	handlers.SetupRoutes(router)
}
