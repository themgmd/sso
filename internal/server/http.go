package server

import (
	"context"
	"log"
	"net/http"
	"sso/internal/config"
)

type Http struct {
	httpServer *http.Server
}

func NewHttpServer(cfg config.HTTPConfig, handler http.Handler) *Http {
	return &Http{
		httpServer: &http.Server{
			Addr:         cfg.GetAddress(),
			Handler:      handler,
			WriteTimeout: cfg.WriteTimeout,
			ReadTimeout:  cfg.ReadTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
	}
}

func (h *Http) MustStart() {
	err := h.httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("http server failed: %s", err.Error())
	}
}

func (h *Http) Stop(ctx context.Context) error {
	return h.httpServer.Shutdown(ctx)
}
