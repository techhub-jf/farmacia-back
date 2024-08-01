package repositories

import (
	"context"
	"fmt"
	"strings"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

func (r *ProductRepository) ListAll(ctx context.Context, pagination dto.Pagination, filterSearch string) ([]entity.Product, int, error) {
	const (
		operation = "Repository.ProductRepository.ListAllProducts"
	)

	offset := pagination.ItemsPerPage * (pagination.Page - 1)
	args := []interface{}{pagination.ItemsPerPage, offset}

	query := `SELECT 
				count(*) OVER() AS total_count,
				id,
				reference,
				brand,
				description,
				type_id,
				stock
		FROM product
		WHERE deleted_at IS NULL `

	if filterSearch != "" {
		query += "AND ("
		if strings.Contains(filterSearch, "#") {
			query += fmt.Sprintf(`reference ILIKE '%%%s%%' OR `, strings.ReplaceAll(filterSearch, "#", ""))
		}

		query += fmt.Sprintf(`description ILIKE '%%%s%%'
			OR brand ILIKE '%%%s%%'
			) `, filterSearch, filterSearch)
	}

	query += fmt.Sprintf(`ORDER BY %s %s
		LIMIT $1 OFFSET $2;	
	`, pagination.SortBy, pagination.SortType)

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
			&product.Brand,
			&product.Description,
			&product.TypeID,
			&product.Stock,
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
