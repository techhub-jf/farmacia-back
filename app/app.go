package app

import (
	"github.com/techhub-jf/farmacia-back/app/config"
	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
	"github.com/techhub-jf/farmacia-back/app/gateway/postgres"
	"github.com/techhub-jf/farmacia-back/app/gateway/postgres/repositories"
)

type App struct {
	UseCase *usecase.UseCase
}

func New(config config.Config, db *postgres.Client) (*App, error) {
	useCase := &usecase.UseCase{
		AppName:              config.App.Name,
		AccountsRepository:   repositories.NewAccountsRepository(db),
		ClientsRepository:    repositories.NewClientsRepository(db),
		DeliveriesRepository: repositories.NewDeliveriesRepository(db),
		MedicinesRepository:  repositories.NewMedicinesRepository(db),
	}

	return &App{
		UseCase: useCase,
	}, nil
}
