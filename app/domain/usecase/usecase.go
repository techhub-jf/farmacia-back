package usecase

import (
	"context"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

type UseCase struct {
	AppName string

	AccountsRepository   accountsRepository
	ClientsRepository    clientsRepository
	DeliveriesRepository deliveriesRepository
	ProductsRepository   productsRepository
}

type accountsRepository interface {
	GetAccountByEmail(ctx context.Context, email string) (entity.Account, error)
}

type deliveriesRepository interface {
	ListAll(ctx context.Context, filters dto.Pagination) ([]entity.Delivery, int, error)
}

type clientsRepository interface {
	GetClients(ctx context.Context, cqp schema.ClientQueryParams) ([]*entity.Client, error)
}
type productsRepository interface {
	ListAll(ctx context.Context, pagination dto.Pagination, filter string) ([]entity.Product, int, error)
}
