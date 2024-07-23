package schema

import (
	"errors"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type ListProductsRequest struct {
	Pagination ListProductsRequestPagination
	Search     string
}
type ListProductsRequestPagination = dto.Pagination

var validSortProductFields = map[string]bool{
	"reference": true,
	"qty":       true,
	"id":        true,
}

func ValidateListProductsRequest(input ListProductsRequestPagination) error {
	if input.Page < 1 {
		return errors.New("page must be greater than 0")
	}

	if input.ItemsPerPage < 1 || input.ItemsPerPage > 100 {
		return errors.New("itemsPerPage must be between 1 and 100")
	}

	if _, found := validSortProductFields[input.SortBy]; !found {
		return errors.New("invalid sortBy param")
	}

	if input.SortType != "ASC" && input.SortType != "DESC" {
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
			UnitID:      product.UnitID,
		})
	}

	return parsedProducts
}
