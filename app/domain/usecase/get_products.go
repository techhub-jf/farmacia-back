package usecase

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type GetProductsInput struct {
	Pagination dto.Pagination
	Search     string
}

type GetProductsOutput struct {
	Products      []entity.Product
	TotalProducts int
}

func (u *UseCase) GetProducts(ctx context.Context, input GetProductsInput) (GetProductsOutput, error) {
	products, totalRecords, err := u.ProductsRepository.ListAll(ctx, input.Pagination, input.Search)
	if err != nil {
		return GetProductsOutput{}, fmt.Errorf("error listing products: %w", err)
	}

	return GetProductsOutput{
		Products:      products,
		TotalProducts: totalRecords,
	}, nil
}
