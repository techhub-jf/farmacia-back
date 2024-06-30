package repositories

import "github.com/techhub-jf/farmacia-back/app/gateway/postgres"

type ClientsRepository struct {
	*postgres.Client
}

func NewClientsRepository(client *postgres.Client) *ClientsRepository {
	return &ClientsRepository{client}
}
