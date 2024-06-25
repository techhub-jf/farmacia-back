package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

const getClientsClause = `
SELECT
id,
reference,
full_name,
cpf,
rg,
phone,
created_at
FROM client
WHERE deleted_at IS NULL
`

func (r ClientsRepository) GetClients(ctx context.Context) ([]*entity.Client, error) {
	const operation = "Repository.ClientsRepository.GetClients"

	rows, _ := r.Pool.Query(ctx, getClientsClause)
	clients, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[entity.Client])
	if err != nil {
		return []*entity.Client{}, fmt.Errorf("%s -> %w", operation, err)
	}

	return clients, nil
}
