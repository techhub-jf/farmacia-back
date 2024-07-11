package repositories

import "github.com/techhub-jf/farmacia-back/app/gateway/postgres"

type DeliveriesRepository struct {
	*postgres.Client
}

func NewDeliveriesRepository(client *postgres.Client) *DeliveriesRepository {
	return &DeliveriesRepository{client}
}
