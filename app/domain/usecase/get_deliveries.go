package usecase

import (
	"context"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/gateway/postgres/repositories"
)

type GetDeliveriesOutput struct {
	Deliveries []*entity.Delivery
	Meta       string
}

type GetDeliveriesInput struct {
	Page     int
	SortBy   string
	SortType string
}

func (u *UseCase) GetDeliveries(ctx context.Context, input GetDeliveriesInput) (GetDeliveriesOutput, error) {
	args := repositories.PaginationFilters{
		Page:      int32(input.Page),
		Sort_by:   input.SortBy,
		Sort_type: input.SortType,
	}
	deliveries, err := u.DeliveriesRepository.GetAll(ctx, args)
	if err != nil {
		return GetDeliveriesOutput{}, err
	}

	return GetDeliveriesOutput{
		Deliveries: deliveries,
		Meta:       "teste",
	}, nil
}
