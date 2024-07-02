package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

type GetDeliveriesInput struct {
	Pagination dto.Pagination
}

type GetDeliveriesOutput struct {
	schema.ListDeliveriesOutput
}

func (u *UseCase) GetDeliveries(ctx context.Context, input GetDeliveriesInput) (GetDeliveriesOutput, error) {
	args := dto.Pagination{
		Page:         input.Pagination.Page,
		SortBy:       input.Pagination.SortBy,
		SortType:     input.Pagination.SortType,
		ItemsPerPage: input.Pagination.ItemsPerPage,
	}

	deliveries, totalRecords, err := u.DeliveriesRepository.ListAll(ctx, args)
	if err != nil {
		return GetDeliveriesOutput{}, fmt.Errorf("error listing deliveries: %w", err)
	}

	metadata := schema.Meta{
		ItemsPerPage: input.Pagination.ItemsPerPage,
		CurrentPage:  input.Pagination.Page,
		TotalItems:   totalRecords,
	}

	return GetDeliveriesOutput{
		schema.ListDeliveriesOutput{
			Items:    deliveries,
			Metadata: metadata,
		},
	}, nil
}
