package usecase

import (
	"context"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type UseCase struct {
	AppName string

	//Accounts
	AccountsRepository  accountsRepository
	MedicinesRepository medicinesRepository
}

type accountsRepository interface {
	GetAccountByEmail(ctx context.Context, email string) (entity.Account, error)
}
type medicinesRepository interface {
	GetMedicines(ctx context.Context) ([]entity.Medicine, error)
}
