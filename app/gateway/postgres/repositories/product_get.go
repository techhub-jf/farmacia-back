package repositories

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

func (r *ProductRepository) ListAll(ctx context.Context, filters dto.Pagination) ([]entity.Product, int, error) {
	const (
		operation = "Repository.MedicinesRepository.GetMedicines"
	)

	offset := filters.ItemsPerPage * (filters.Page - 1)
	args := []interface{}{filters.ItemsPerPage, offset}

	query := fmt.Sprintf(`
		SELECT 
				count(*) OVER() AS total_count,
				id,
				reference,
				active_principle,
				description,
				unit_id,
				qty
		FROM product
		WHERE deleted_at IS NULL
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
	products := []entity.Product{}

	for rows.Next() {
		var product entity.Product

		err := rows.Scan(
			&totalRecords,
			&product.ID,
			&product.Reference,
			&product.Description,
			&product.ActivePrinciple,
			&product.UnitID,
			&product.Qty,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("%s: %w", operation, err)
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("%s -> %w", operation, err)
	}

	return products, totalRecords, nil
}
