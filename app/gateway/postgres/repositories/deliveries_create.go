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

	query := `
		INSERT INTO deliveries (reference,  client_id,  medicine_id, qty, unit_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, reference, created_at;
	`

	args := []interface{}{
		input.Delivery.Reference, input.Delivery.ClientID,
		input.Delivery.MedicineID, input.Delivery.Qty, input.Delivery.UnitID,
	}

	row := r.Client.Pool.QueryRow(
		ctx,
		query,
		args...,
	)

	delivery := entity.Delivery{}
	err := row.Scan(&delivery.ID,
		&delivery.Reference, &delivery.CreatedAt)
	if err != nil {
		return entity.Delivery{}, fmt.Errorf("%s: %w", operation, err)
	}

	return delivery, nil
}
