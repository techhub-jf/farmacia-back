package usecase

import (
	"context"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type UseCase struct {
	AppName string

	AccountsRepository accountsRepository
}

type accountsRepository interface {
	GetAccountByEmail(ctx context.Context, email string) (entity.Account, error)
}
