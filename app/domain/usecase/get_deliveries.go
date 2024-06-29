package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

func (u *UseCase) GetDeliveries(ctx context.Context, input schema.ListDeliveriesInput) (schema.ListDeliveriesOutput, error) {
	args := schema.Pagination{
		Page:         input.Page,
		SortBy:       input.SortBy,
		SortType:     input.SortType,
		ItemsPerPage: input.ItemsPerPage,
	}

	deliveries, totalRecords, err := u.DeliveriesRepository.GetAll(ctx, args)
	if err != nil {
		return schema.ListDeliveriesOutput{}, fmt.Errorf("error listing deliveries: %w", err)
	}

	metadata := schema.Meta{
		ItemsPerPage: input.ItemsPerPage,
		CurrentPage:  input.Page,
		TotalItems:   totalRecords,
	}

	return schema.ListDeliveriesOutput{
		Items:    deliveries,
		Metadata: metadata,
	}, nil
}
