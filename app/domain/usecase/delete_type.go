package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type DeleteTypeInput struct {
	ID int32
}

type DeleteTypeOutput struct {
	Delivery entity.Delivery
}

func (u *UseCase) DeleteType(ctx context.Context, input DeleteTypeInput) error {
	err := u.TypeRepository.Delete(ctx, input.ID)
	if err != nil {
		return fmt.Errorf("error deleting delivery: %w", err)
	}

	return nil
}
