package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

func (r *TypeRepository) ListAll(ctx context.Context, pagination dto.Pagination, filterSearch string) ([]entity.Type, int, error) {
	const (
		operation = "Repository.TypeRepository.ListAllTypes"
	)

	offset := pagination.ItemsPerPage * (pagination.Page - 1)
	args := []interface{}{pagination.ItemsPerPage, offset}

	query := `SELECT 
				count(*) OVER() AS total_count,
				id,
				label,
		FROM type
		WHERE deleted_at IS NULL `

	if filterSearch != "" {
		query += "AND ("
		if strings.Contains(filterSearch, "#") {
			query += fmt.Sprintf(`reference ILIKE '%%%s%%' OR `, strings.ReplaceAll(filterSearch, "#", ""))
		}

		query += fmt.Sprintf(`label ILIKE '%%%s%%') `, filterSearch)
	}

	query += fmt.Sprintf(`ORDER BY %s %s
		LIMIT $1 OFFSET $2;	
	`, pagination.SortBy, pagination.SortType)

	rows, err := r.Client.Pool.Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("%s: %w", operation, err)
	}
	defer rows.Close()

	totalRecords := 0
	types := []entity.Type{}

	for rows.Next() {
		var t entity.Type

		err := rows.Scan(
			&totalRecords,
			&t.ID,
			&t.Reference,
			&t.Label,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("%s: %w", operation, err)
		}

		types = append(types, t)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%s -> %w", operation, err)
	}

	return types, totalRecords, nil
}
