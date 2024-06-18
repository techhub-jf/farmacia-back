package app

import (
	"github.com/techhub-jf/farmacia-back/app/config"
	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
	"github.com/techhub-jf/farmacia-back/app/gateway/postgres"
	"github.com/techhub-jf/farmacia-back/app/gateway/postgres/respositories"
)

type App struct {
	UseCase *usecase.UseCase
}

func New(config config.Config, db *postgres.Client) (*App, error) { //nolint: revive

	useCase := &usecase.UseCase{
		AppName:             config.App.Name,
		AccountsRepository:  respositories.NewAccountsRepository(db),
		MedicinesRepository: respositories.NewMedicinesRepository(db),
	}

	return &App{
		UseCase: useCase,
	}, nil
}
