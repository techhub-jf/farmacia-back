package repositories

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

func (r *DeliveriesRepository) ListAll(ctx context.Context, filters dto.Pagination) ([]entity.Delivery, int, error) {
	const (
		operation = "Repository.DeliveriesRepository.ListAll"
	)

	offset := filters.ItemsPerPage * (filters.Page - 1)
	args := []interface{}{filters.ItemsPerPage, offset}

	query := fmt.Sprintf(`
		SELECT 
				count(*) OVER() AS total_count,
				id,
				reference,
				qty,
				created_at
		FROM deliveries
		ORDER BY %s %s
		LIMIT $1 OFFSET $2;	
	`, filters.SortBy, filters.SortType)

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

	var deliveries []entity.Delivery

	for rows.Next() {
		var delivery entity.Delivery

		err := rows.Scan(
			&totalRecords,
			&delivery.ID,
			&delivery.Reference,
			&delivery.Qty,
			&delivery.CreatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("%s: %w", operation, err)
		}

		deliveries = append(deliveries, delivery)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%s -> %w", operation, err)
	}

	return deliveries, totalRecords, nil
}

func (r *DeliveriesRepository) GetByReference(ctx context.Context, reference string) (entity.Delivery, error) {
	const (
		operation = "Repository.DeliveriesRepository.GetByReference"
	)

	query := `
		SELECT 
			id,
			reference,
			qty,
			created_at,
			updated_at
		FROM deliveries
		WHERE reference = $1;
	`

	var delivery entity.Delivery

	err := r.Client.Pool.QueryRow(ctx, query, reference).Scan(
		&delivery.ID,
		&delivery.Reference,
		&delivery.Qty,
		&delivery.CreatedAt,
		&delivery.UpdatedAt,
	)
	if err != nil {
		return entity.Delivery{}, fmt.Errorf("%s: %w", operation, err)
	}

	return delivery, nil
}
