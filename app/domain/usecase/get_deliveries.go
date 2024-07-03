package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type GetDeliveriesInput struct {
	Pagination dto.Pagination
}

type GetDeliveriesOutput struct {
	Deliveries      []entity.Delivery
	TotalDeliveries int
}

func (u *UseCase) GetDeliveries(ctx context.Context, input GetDeliveriesInput) (GetDeliveriesOutput, error) {
	deliveries, totalRecords, err := u.DeliveriesRepository.ListAll(ctx, input.Pagination)
	if err != nil {
		return GetDeliveriesOutput{}, fmt.Errorf("error listing deliveries: %w", err)
	}

	return GetDeliveriesOutput{
		Deliveries:      deliveries,
		TotalDeliveries: totalRecords,
	}, nil
}
