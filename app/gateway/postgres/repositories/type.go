package repositories

import "github.com/techhub-jf/farmacia-back/app/gateway/postgres"

type TypeRepository struct {
	*postgres.Client
}

func NewTypeRepository(client *postgres.Client) *TypeRepository {
	return &TypeRepository{client}
}
