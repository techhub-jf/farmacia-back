package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type DeleteDeliveryInput struct {
	Id int32
}

type DeleteDeliveryOutput struct {
	Delivery entity.Delivery
}

func (u *UseCase) DeleteDelivery(ctx context.Context, input DeleteDeliveryInput) error {
	delivery, err := u.DeliveriesRepository.GetByID(ctx, input.Id)
	if err != nil {
		return fmt.Errorf("error getting delivery: %w", err)
	}

	if delivery.DeletedAt != nil {
		return fmt.Errorf("delivery already deleted")
	}

	err = u.DeliveriesRepository.Delete(ctx, input.Id)
	if err != nil {
		return fmt.Errorf("error deleting delivery: %w", err)
	}

	return nil
}
