package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type DeleteDeliveryInput struct {
	ID int32
}

type DeleteDeliveryOutput struct {
	Delivery entity.Delivery
}

func (u *UseCase) DeleteDelivery(ctx context.Context, input DeleteDeliveryInput) error {
	err := u.DeliveriesRepository.Delete(ctx, input.ID)
	if err != nil {
		return fmt.Errorf("error deleting delivery: %w", err)
	}
	return nil
}
