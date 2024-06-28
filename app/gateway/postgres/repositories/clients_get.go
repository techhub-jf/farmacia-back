package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
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
ORDER BY %s %s
LIMIT $1 OFFSET $2
`

func (r *ClientsRepository) GetClients(
	ctx context.Context, cqp schema.ValidatedClientQueryParams) (
	[]*entity.Client, error,
) {
	const operation = "Repository.ClientsRepository.GetClients"

	finalQuery := fmt.Sprintf(getClientsClause, cqp.SortBy, cqp.SortType)

	offset := (cqp.Page - 1) * cqp.Limit
	rows, _ := r.Pool.Query(ctx, finalQuery, cqp.Limit, offset)

	clients, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[entity.Client])
	if err != nil {
		return []*entity.Client{}, fmt.Errorf("%s -> %w", operation, err)
	}

	return clients, nil
}
