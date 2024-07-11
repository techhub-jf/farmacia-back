package repositories

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

func (r *MedicinesRepository) GetMedicines(ctx context.Context) ([]entity.Medicine, error) {
	const (
		operation           = "Repository.MedicinesRepository.GetMedicines"
		getAllMedicinesStmt = `SELECT id, reference, client_id, medicine_id, qty, unit_id, created_at, updated_at, deleted_at FROM medicine`
	)

	rows, err := r.Client.Pool.Query(ctx, getAllMedicinesStmt)
	if err != nil {
		return nil, fmt.Errorf("%s -> %w", operation, err)
	}
	defer rows.Close()

	var medicines []entity.Medicine
	for rows.Next() {
		var medicine entity.Medicine
		if err := rows.Scan(
			&medicine.ID,
			&medicine.Reference,
			&medicine.Client_id,
			&medicine.Medicine_id,
			&medicine.Qty,
			&medicine.Unit_id,
			&medicine.CreatedAt,
			&medicine.UpdatedAt,
			&medicine.DeletedAt,
		); err != nil {
			return nil, fmt.Errorf("%s -> %w", operation, err)
		}
		medicines = append(medicines, medicine)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("%s -> %w", operation, rows.Err())
	}

	return medicines, nil
}
