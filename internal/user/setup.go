package user

import (
	"github.com/gorilla/mux"
	userPostgre "sso/internal/user/postgre"
	"sso/pkg/connectors/postgre"
)

func Setup(db *postgre.DB, _ *mux.Router) *User {
	userRepo := userPostgre.NewUser(db)
	userService := New(userRepo)
	//userHandler := dhttp.NewUser(userService)
	//userHandler.SetupRoutes(router)

	return userService
}
