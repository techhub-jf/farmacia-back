package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type GetTypeByLabelInput struct {
	Label string
}

type GetTypeByLabelOutput struct {
	Type entity.Type
}

func (u *UseCase) GetTypeByLabel(ctx context.Context, input GetTypeByLabelInput) (GetTypeByLabelOutput, error) {
	t, err := u.TypeRepository.GetByLabel(ctx, input.Label)
	if err != nil {
		return GetTypeByLabelOutput{}, fmt.Errorf("error getting type: %w", err)
	}

	return GetTypeByLabelOutput{
		Type: t,
	}, nil
}
