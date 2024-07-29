package usecase

import (
	"context"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

type UseCase struct {
	AppName string

	AccountsRepository accountsRepository
	ClientsRepository  clientsRepository
}

type accountsRepository interface {
	GetAccountByEmail(ctx context.Context, email string) (entity.Account, error)
}

type clientsRepository interface {
	GetClients(ctx context.Context, cqp schema.ClientQueryParams) ([]*entity.Client, error)
	CreateClient(ctx context.Context, client entity.Client) (entity.Client, error)
	UpdateClient(ctx context.Context, client entity.Client) (entity.Client, error)
}
