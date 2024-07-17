package schema

import (
	"errors"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type ListProductsRequest = dto.Pagination

var validSearchProductFields = map[string]bool{
	"reference": true,
	"qty":       true,
	"id":        true,
}

func ValidateListProductsRequest(input ListProductsRequest) error {
	if input.Page < 1 {
		return errors.New("page must be greater than 0")
	}

	if input.ItemsPerPage < 1 || input.ItemsPerPage > 100 {
		return errors.New("itemsPerPage must be between 1 and 100")
	}

	if _, found := validSearchProductFields[input.SortBy]; !found {
		return errors.New("invalid sortBy param")
	}

	if input.SortType != "ASC" && input.SortBy != "DESC" {
		return errors.New("sortType param must be 'ASC' or 'DESC'")
	}

	return nil
}

type ListProductsResponse struct {
	ID              uint   `json:"id"`
	Reference       string `json:"reference"`
	Qty             uint   `json:"qty"`
	Description     string `json:"description"`
	ActivePrinciple string `json:"active_principle"`
	UnitID          uint   `json:"unit_id"`
}

type ListProductsOutput = PaginatedResponse[ListProductsResponse]

func ConvertProductsToListResponse(products []entity.Product) []ListProductsResponse {
	parsedProducts := []ListProductsResponse{}

	for _, product := range products {
		parsedProducts = append(parsedProducts, ListProductsResponse{
			ID:              product.ID,
			Reference:       product.Reference,
			Description:     product.Description,
			ActivePrinciple: product.ActivePrinciple,
			Qty:             product.Qty,
		})
	}

	return parsedProducts
}
