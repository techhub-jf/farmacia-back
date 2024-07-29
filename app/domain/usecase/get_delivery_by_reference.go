package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type GetDeliveryByReferenceInput struct {
	Reference string
}

type GetDeliveryByReferenceOutput struct {
	Delivery entity.Delivery
}

func (u *UseCase) GetDeliveryByReference(ctx context.Context, input GetDeliveryByReferenceInput) (GetDeliveryByReferenceOutput, error) {
	delivery, err := u.DeliveriesRepository.GetByReference(ctx, input.Reference)
	if err != nil {
		return GetDeliveryByReferenceOutput{}, fmt.Errorf("error getting delivery: %w", err)
	}

	return GetDeliveryByReferenceOutput{
		Delivery: delivery,
	}, nil
}
