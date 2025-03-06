package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/techhub-jf/farmacia-back/app/config"
	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
)

type Handler struct {
	cfg     config.Config
	useCase *usecase.UseCase
}

func New(cfg config.Config, useCase *usecase.UseCase) Handler {
	return Handler{
		cfg:     cfg,
		useCase: useCase,
	}
}

func RegisterHealthCheckRoute(router chi.Router) {
	router.Get("/healthcheck", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
}

func RegisterPublicRoutes(
	router chi.Router,
	cfg config.Config,
	useCase *usecase.UseCase,
) {
	handler := New(cfg, useCase)
	handler.LoginSetup(router)
	handler.ListClients(router)
}

func RegisterPrivateRoutes(
	router chi.Router,
	cfg config.Config,
	useCase *usecase.UseCase,
) {
	handler := New(cfg, useCase)
	handler.DeliveriesSetup(router)
	handler.ProductsSetup(router)
	handler.TypesSetup(router)
}
