package schema

import (
	"errors"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type ListProductsRequest struct {
	Pagination dto.Pagination
	Search     string
}

var validSortProductFields = map[string]bool{
	"reference": true,
	"qty":       true,
	"id":        true,
}

func ValidateListProductsRequest(input ListProductsRequest) error {
	if input.Pagination.Page < 1 {
		return errors.New("page must be greater than 0")
	}

	if input.Pagination.ItemsPerPage < 1 || input.Pagination.ItemsPerPage > 100 {
		return errors.New("itemsPerPage must be between 1 and 100")
	}

	if _, found := validSortProductFields[input.Pagination.SortBy]; !found {
		return errors.New("invalid sortBy param")
	}

	if input.Pagination.SortType != "ASC" && input.Pagination.SortType != "DESC" {
		return errors.New("sortType param must be 'ASC' or 'DESC'")
	}

	return nil
}

type ListProductsResponse struct {
	ID          uint   `json:"id"`
	Reference   string `json:"reference"`
	Stock       uint   `json:"stock"`
	Description string `json:"description"`
	Branch      string `json:"branch"`
	UnitID      uint   `json:"unit_id"`
}

type ListProductsOutput = PaginatedResponse[ListProductsResponse]

func ConvertProductsToListResponse(products []entity.Product) []ListProductsResponse {
	parsedProducts := []ListProductsResponse{}

	for _, product := range products {
		parsedProducts = append(parsedProducts, ListProductsResponse{
			ID:          product.ID,
			Reference:   product.Reference,
			Description: product.Description,
			Branch:      product.Branch,
			Stock:       product.Stock,
		})
	}

	return parsedProducts
}
