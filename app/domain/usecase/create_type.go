package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type CreateTypeInput struct {
	Type dto.Type
}

type CreateTypeOutput struct {
	Type entity.Type
}

func (u *UseCase) CreateType(ctx context.Context, input CreateTypeInput) (CreateTypeOutput, error) {
	t, err := u.TypeRepository.Create(ctx, input)
	if err != nil {
		return CreateTypeOutput{}, fmt.Errorf("error creating delivery: %w", err)
	}

	return CreateTypeOutput{
		Type: t,
	}, nil
}
