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
	TypeRepository       typeRepository
}

type accountsRepository interface {
	GetAccountByEmail(ctx context.Context, email string) (entity.Account, error)
}

type deliveriesRepository interface {
	ListAll(ctx context.Context, filters dto.Pagination) ([]entity.Delivery, int, error)
	GetByID(ctx context.Context, id int32) (entity.Delivery, error)
	GetByReference(ctx context.Context, reference string) (entity.Delivery, error)
	Create(ctx context.Context, delivery CreateDeliveryInput) (entity.Delivery, error)
	Delete(ctx context.Context, id int32) error
}

type clientsRepository interface {
	GetClients(ctx context.Context, cqp schema.ClientQueryParams) ([]*entity.Client, error)
}
type productsRepository interface {
	ListAll(ctx context.Context, pagination dto.Pagination, filter string) ([]entity.Product, int, error)
}
type typeRepository interface {
	ListAll(ctx context.Context, pagination dto.Pagination, filter string) ([]entity.Type, int, error)
	GetByReference(ctx context.Context, reference string) (entity.Type, error)
	GetByLabel(ctx context.Context, label string) (entity.Type, error)
	Create(ctx context.Context, t CreateTypeInput) (entity.Type, error)
	Update(ctx context.Context, t UpdateTypeInput) (entity.Type, error)
	Delete(ctx context.Context, id int32) error
}
