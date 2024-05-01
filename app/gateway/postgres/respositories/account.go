package respositories

import "github.com/techhub-jf/farmacia-back/app/gateway/postgres"

type AccountsRepository struct {
	*postgres.Client
}

func NewAccountsRepository(client *postgres.Client) *AccountsRepository {
	return &AccountsRepository{client}
}
