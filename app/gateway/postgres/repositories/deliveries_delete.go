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
		SET deleted_at = COALESCE(deleted_at, current_timestamp)
		WHERE id = $1;
	`
	res, err := r.Client.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("%s: %s", operation, "no rows in result set")
	}
	return nil
}
