package repositories

import (
	"context"
	"fmt"
)

func (r *DeliveriesRepository) Delete(ctx context.Context, id int32) error {
	const (
		operation = "Repository.DeliveriesRepository.Delete"
	)

	query := `
		UPDATE deliveries
		SET deleted_at = current_timestamp
		WHERE id = $1;
	`

	_, err := r.Client.Pool.Exec(ctx, query, id)

	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
