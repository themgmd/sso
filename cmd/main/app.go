package main

import (
	"context"
	"log"
	"os/signal"
	"sso/internal/config"
	"sso/internal/core"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	config.Init()

	app := core.New()
	err := app.Run(ctx)
	if err != nil {
		log.Fatalf("error occurred while app started: %s", err.Error())
	}

	log.Printf("App started on port: %s", config.Get().HTTP.Port)
	<-ctx.Done()

	err = app.Stop(ctx)
	if err != nil {
		log.Fatalf("error occured while app stopped: %s", err.Error())
	}
}
