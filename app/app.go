package app

import (
	"github.com/techhub-jf/farmacia-back/app/config"
	"github.com/techhub-jf/farmacia-back/app/gateway/postgres"
)

type App struct {
}

func New(config config.Config, db *postgres.Client) (*App, error) { //nolint: revive

	return &App{}, nil
}
