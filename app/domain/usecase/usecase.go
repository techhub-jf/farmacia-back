package usecase

import (
	"context"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type UseCase struct {
	AppName string

	//Accounts
	AccountsRepository accountsRepository
}

type accountsRepository interface {
	GetAccountByEmail(ctx context.Context, cpf string) (entity.Account, error)
}
