package repositories

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
)

func (r *TypeRepository) Create(ctx context.Context, input usecase.CreateTypeInput) (entity.Type, error) {
	const (
		operation = "Repository.DeliveriesRepository.Create"
	)

	query := `
		INSERT INTO type (reference, label)
		VALUES ($1, $2)
		RETURNING id, reference, created_at;
	`

	args := []interface{}{
		input.Type.Reference,
		input.Type.Label,
	}

	row := r.Client.Pool.QueryRow(
		ctx,
		query,
		args...,
	)

	t := entity.Type{}

	err := row.Scan(&t.ID,
		&t.Reference, &t.CreatedAt)
	if err != nil {
		return entity.Type{}, fmt.Errorf("%s: %w", operation, err)
	}

	return t, nil
}
