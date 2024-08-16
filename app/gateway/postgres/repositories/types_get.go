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
				reference,
				label,
				created_at
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
		var typeStruct entity.Type

		err := rows.Scan(
			&totalRecords,
			&typeStruct.ID,
			&typeStruct.Reference,
			&typeStruct.Label,
			&typeStruct.CreatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("%s: %w", operation, err)
		}

		types = append(types, typeStruct)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%s -> %w", operation, err)
	}

	return types, totalRecords, nil
}

func (r *TypeRepository) GetByLabel(ctx context.Context, label string) (entity.Type, error) {
	const (
		operation = "Repository.TypeRepository.GetByLabel"
	)

	query := `
		SELECT 
			id,
			reference,
			label,
			created_at,
			updated_at
		FROM type
		WHERE label = $1;
	`

	var typeStruct entity.Type

	err := r.Client.Pool.QueryRow(ctx, query, label).Scan(
		&typeStruct.ID,
		&typeStruct.Reference,
		&typeStruct.Label,
		&typeStruct.CreatedAt,
		&typeStruct.UpdatedAt,
	)
	if err != nil {
		return entity.Type{}, fmt.Errorf("%s: %w", operation, err)
	}

	return typeStruct, nil
}

func (r *TypeRepository) GetByReference(ctx context.Context, reference string) (entity.Type, error) {
	const (
		operation = "Repository.TypeRepository.GetByReference"
	)

	query := `
		SELECT 
			id,
			reference,
			label,
			created_at,
			updated_at
		FROM type
		WHERE reference = $1;
	`

	var typeStruct entity.Type

	err := r.Client.Pool.QueryRow(ctx, query, reference).Scan(
		&typeStruct.ID,
		&typeStruct.Reference,
		&typeStruct.Label,
		&typeStruct.CreatedAt,
		&typeStruct.UpdatedAt,
	)
	if err != nil {
		return entity.Type{}, fmt.Errorf("%s: %w", operation, err)
	}

	return typeStruct, nil
}
