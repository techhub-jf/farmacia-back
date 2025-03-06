package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type GetTypesInput struct {
	Pagination dto.Pagination
	Search     string
}

type GetTypesOutput struct {
	Types      []entity.Type
	TotalTypes int
}

func (u *UseCase) GetTypes(ctx context.Context, input GetTypesInput) (GetTypesOutput, error) {
	types, totalRecords, err := u.TypeRepository.ListAll(ctx, input.Pagination, input.Search)
	if err != nil {
		return GetTypesOutput{}, fmt.Errorf("error listing types: %w", err)
	}

	return GetTypesOutput{
		Types:      types,
		TotalTypes: totalRecords,
	}, nil
}
