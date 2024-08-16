package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type UpdateTypeInput struct {
	Type dto.Type
}

type UpdateTypeOutput struct {
	Type entity.Type
}

func (u *UseCase) UpdateType(ctx context.Context, input UpdateTypeInput) (UpdateTypeOutput, error) {
	output, err := u.TypeRepository.Update(ctx, input)
	if err != nil {
		return UpdateTypeOutput{}, fmt.Errorf("error getting type: %w", err)
	}

	return UpdateTypeOutput{
		Type: output,
	}, nil
}
