package repositories

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

const defaultLimit = 10

type PaginationFilters struct {
	Page      int32
	Sort_by   string
	Sort_type string
}

func (r *DeliveriesRepository) GetAll(ctx context.Context, filters PaginationFilters) ([]*entity.Delivery, error) {
	const (
		operation = "Repository.DeliveriesRepository.GetDeliveries"
	)

	offset := defaultLimit * (filters.Page - 1)
	args := []interface{}{defaultLimit, offset}

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
	`, filters.Sort_by, filters.Sort_type)

	rows, err := r.Client.Pool.Query(
		ctx,
		query,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	totalRecords := 0
	deliveries := []*entity.Delivery{}

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
			return nil, err
		}
		deliveries = append(deliveries, &delivery)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s -> %w", operation, err)
	}

	print("Total: ", totalRecords)

	return deliveries, nil
}
