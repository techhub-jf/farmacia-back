package repositories

import "github.com/techhub-jf/farmacia-back/app/gateway/postgres"

type ProductRepository struct {
	*postgres.Client
}

func NewProductRepository(client *postgres.Client) *ProductRepository {
	return &ProductRepository{client}
}
