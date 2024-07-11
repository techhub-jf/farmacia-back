package api

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/techhub-jf/farmacia-back/app/config"
	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/middleware"
)

type API struct {
	Handler http.Handler
	cfg     config.Config
	useCase *usecase.UseCase
}

func New(cfg config.Config, useCase *usecase.UseCase) *API {
	api := &API{
		cfg:     cfg,
		useCase: useCase,
	}

	api.setupRouter()

	return api
}

func (api *API) setupRouter() {
	router := chi.NewRouter()

	router.Use(middleware.CORS)

	api.registerRoutes(router)

	api.Handler = router
}

func (api *API) registerRoutes(router *chi.Mux) {
	handler.RegisterHealthCheckRoute(router)
	router.Route("/api/v1/farmacia-tech", func(publicRouter chi.Router) {
		handler.RegisterPublicRoutes(
			publicRouter,
			api.cfg,
			api.useCase,
		)

		publicRouter.Group(func(privateRouter chi.Router) {
			privateRouter.Use(middleware.ProtectedHandler(api.cfg.JwtSecretKey))
			handler.RegisterPrivateRoutes(
				privateRouter,
				api.cfg,
				api.useCase,
			)
		})
	})
}
