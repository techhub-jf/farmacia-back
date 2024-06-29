package repositories

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

func (r *DeliveriesRepository) GetAll(ctx context.Context, filters dto.Pagination) ([]*schema.ListDeliveriesResponse, int, error) {
	const (
		operation = "Repository.DeliveriesRepository.GetAll"
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
	deliveries := []*schema.ListDeliveriesResponse{}

	for rows.Next() {
		var delivery schema.ListDeliveriesResponse

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

		deliveries = append(deliveries, &delivery)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%s -> %w", operation, err)
	}

	return deliveries, totalRecords, nil
}
