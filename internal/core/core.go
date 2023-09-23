package core

import (
	"context"
	"sso/internal/auth"
	"sso/internal/config"
	"sso/internal/server"
	"sso/internal/server/middleware"
	"sso/internal/user"
	"sso/migration"
	"sso/pkg/connectors/postgre"

	"github.com/gorilla/mux"
)

type Core struct {
	httpServer *server.Http
	db         *postgre.DB
}

func New() *Core {
	return &Core{}
}

func (c *Core) Run(_ context.Context) (err error) {
	cfg := config.Get()

	c.db, err = postgre.New(cfg.Postgre)
	if err != nil {
		return err
	}

	_, err = migration.Apply()
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	userService := user.Setup(c.db, router)
	auth.Setup(userService, router)

	manager := middleware.NewManager(router)
	manager.Apply()

	c.httpServer = server.NewHttpServer(cfg.HTTP, router)
	go c.httpServer.MustStart()
	return nil
}

func (c *Core) Stop(ctx context.Context) error {
	err := c.httpServer.Stop(ctx)
	if err != nil {
		return err
	}

	err = c.db.Close()
	if err != nil {
		return err
	}

	return nil
}
