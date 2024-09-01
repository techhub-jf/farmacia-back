package repositories

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
)

func (r *DeliveriesRepository) Create(ctx context.Context, input usecase.CreateDeliveryInput) (entity.Delivery, error) {
	const (
		operation = "Repository.DeliveriesRepository.Create"
	)

	tx, err := r.Client.Pool.Begin(ctx)
	if err != nil {
		return entity.Delivery{}, fmt.Errorf("%s: %w", operation, err)
	}

	defer func() {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil && err == nil {
			err = rollbackErr
		}
	}()

	query := `
		INSERT INTO deliveries (reference, client_id, qty, unit_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, reference, created_at;
	`

	args := []interface{}{
		input.Delivery.Reference, input.Delivery.ClientID,
		input.Delivery.Qty, input.Delivery.UnitID,
	}

	var delivery entity.Delivery

	err = tx.QueryRow(ctx, query, args...).Scan(&delivery.ID, &delivery.Reference, &delivery.CreatedAt)
	if err != nil {
		return entity.Delivery{}, fmt.Errorf("%s: %w", operation, err)
	}

	query = `
		INSERT INTO delivery_product (delivery_id, product_id)
		SELECT $1, unnest($2::int[]);
	`

	_, err = tx.Exec(ctx, query, delivery.ID, input.Delivery.ProductIDs)
	if err != nil {
		return entity.Delivery{}, fmt.Errorf("%s: %w", operation, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return entity.Delivery{}, fmt.Errorf("%s: %w", operation, err)
	}

	return delivery, nil
}
