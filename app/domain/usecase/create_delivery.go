package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type CreateDeliveryInput struct {
	Delivery dto.Delivery
}

type CreateDeliveryOutput struct {
	Delivery entity.Delivery
}

func (u *UseCase) CreateDelivery(ctx context.Context, input CreateDeliveryInput) (CreateDeliveryOutput, error) {
	delivery, err := u.DeliveriesRepository.Create(ctx, input)
	if err != nil {
		return CreateDeliveryOutput{}, fmt.Errorf("error creating delivery: %w", err)
	}

	return CreateDeliveryOutput{
		Delivery: delivery,
	}, nil
}
