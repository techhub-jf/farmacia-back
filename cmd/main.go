package main

import (
	"context"
	"embed"

	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger" // Import the http-swagger middleware
	"github.com/techhub-jf/farmacia-back/app"
	"github.com/techhub-jf/farmacia-back/app/config"
	"github.com/techhub-jf/farmacia-back/app/gateway/api"
	"github.com/techhub-jf/farmacia-back/app/gateway/postgres"
	_ "github.com/techhub-jf/farmacia-back/docs"
	"golang.org/x/sync/errgroup"
)

var swaggerJSON embed.FS

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load configurations: %v", err)
	}

	// Postgres
	postgresClient, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to start postgres: %v", err)
	}

	// Application
	appl, err := app.New(cfg, postgresClient)
	if err != nil {
		log.Fatalf("failed to start application: %v", err)
	}

	// Create a chi router
	r := chi.NewRouter()

	// Set up API routes

	// Mount your API handler onto the router
	api := api.New(cfg, appl.UseCase)
	r.Mount("/", api.Handler)

	r.Get("/swagger/*", httpSwagger.WrapHandler)
	// Server
	server := &http.Server{
		Addr:         cfg.Server.Address,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// Graceful Shutdown
	stopCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	group, groupCtx := errgroup.WithContext(stopCtx)

	group.Go(func() error {
		log.Printf("starting api server")

		return server.ListenAndServe()
	})

	//nolint:contextcheck
	group.Go(func() error {
		<-groupCtx.Done()

		log.Printf("stopping api; interrupt signal received")

		timeoutCtx, cancel := context.WithTimeout(context.Background(), cfg.App.GracefulShutdownTimeout)
		defer cancel()

		var errs error

		if err := server.Shutdown(timeoutCtx); err != nil {
			errs = errors.Join(errs, fmt.Errorf("failed to stop server: %w", err))
		}

		postgresClient.Close()

		return errs
	})

	if err := group.Wait(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("api exit reason: %v", err)
	}

	stop()
}
