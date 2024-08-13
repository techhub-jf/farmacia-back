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
	MedicinesRepository  medicinesRepository
	ProductsRepository   productsRepository
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
type medicinesRepository interface {
	GetMedicines(ctx context.Context) ([]entity.Medicine, error)
}
type productsRepository interface {
	ListAll(ctx context.Context, pagination dto.Pagination, filter string) ([]entity.Product, int, error)
}
