package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type GetTypeByReferenceInput struct {
	Reference string
}

type GetTypeByReferenceOutput struct {
	Type entity.Type
}

func (u *UseCase) GetTypeByReference(ctx context.Context, input GetTypeByReferenceInput) (GetTypeByReferenceOutput, error) {
	t, err := u.TypeRepository.GetByReference(ctx, input.Reference)
	if err != nil {
		return GetTypeByReferenceOutput{}, fmt.Errorf("error getting type: %w", err)
	}

	return GetTypeByReferenceOutput{
		Type: t,
	}, nil
}
