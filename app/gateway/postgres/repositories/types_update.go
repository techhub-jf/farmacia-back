package repositories

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
)

func (r *TypeRepository) Update(ctx context.Context, input usecase.UpdateTypeInput) (entity.Type, error) {
	const (
		operation = "Repository.TypeRepository.Update"
	)

	query := `
		UPDATE type
		SET label = $2
		WHERE id = $1
		RETURNING id, reference, created_at;
	`

	args := []interface{}{
		input.Type.ID,
		input.Type.Label,
	}

	row := r.Client.Pool.QueryRow(
		ctx,
		query,
		args...)

	t := entity.Type{}

	err := row.Scan(&t.ID,
		&t.Reference, &t.CreatedAt)
	if err != nil {
		return entity.Type{}, fmt.Errorf("%s: %w", operation, err)
	}

	return t, nil
}
