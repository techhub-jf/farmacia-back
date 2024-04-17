package api

import (
	"net/http"

	"github.com/techhub-jf/farmacia-back/app/config"
)

type API struct {
	Handler http.Handler
	cfg     config.Config
}

func New(cfg config.Config) *API {
	api := &API{
		cfg: cfg,
	}

	return api
}
