package usecase

import (
	"context"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/gateway/postgres/repositories"
)

type UseCase struct {
	AppName string

	AccountsRepository   accountsRepository
	DeliveriesRepository deliveriesRepository
}

type accountsRepository interface {
	GetAccountByEmail(ctx context.Context, email string) (entity.Account, error)
}

type deliveriesRepository interface {
	GetAll(ctx context.Context, filters repositories.PaginationFilters) ([]*entity.Delivery, error)
}
