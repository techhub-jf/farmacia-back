package schema

import (
	"errors"
	"time"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type ListTypesRequest struct {
	Pagination ListTypesRequestPagination
	Search     string
}
type ListTypesRequestPagination = dto.Pagination

var validSortTypesFields = map[string]bool{
	"reference":  true,
	"label":      true,
	"created_at": true,
}

type CreateTypeRequest struct {
	Label string `json:"label"`
}

type CreatedTypeResponse struct {
	ID        uint      `json:"id"`
	Reference string    `json:"reference"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTypeResponse struct {
	Type CreatedTypeResponse `json:"type"`
}

func ValidateListTypesRequest(input ListTypesRequestPagination) error {
	if input.Page < 1 {
		return errors.New("page must be greater than 0")
	}

	if input.ItemsPerPage < 1 || input.ItemsPerPage > 100 {
		return errors.New("itemsPerPage must be between 1 and 100")
	}

	if _, found := validSortTypesFields[input.SortBy]; !found {
		return errors.New("invalid sortBy param")
	}

	if input.SortType != "ASC" && input.SortType != "DESC" {
		return errors.New("sortType param must be 'ASC' or 'DESC'")
	}

	return nil
}

type ListTypesResponse struct {
	ID        uint      `json:"id"`
	Reference string    `json:"reference"`
	Label     string    `json:"label"`
	CreatedAt time.Time `json:"created_at"`
}

type ListTypesOutput = PaginatedResponse[ListTypesResponse]

func ConvertTypesToListResponse(types []entity.Type) []ListTypesResponse {
	parsedTypes := []ListTypesResponse{}

	for _, t := range types {
		parsedTypes = append(parsedTypes, ListTypesResponse{
			ID:        t.ID,
			Reference: t.Reference,
			Label:     t.Label,
			CreatedAt: t.CreatedAt,
		})
	}

	return parsedTypes
}

func ConvertTypeToCreateResponse(t entity.Type) CreatedTypeResponse {
	return CreatedTypeResponse{
		ID:        t.ID,
		Reference: t.Reference,
		CreatedAt: t.CreatedAt,
	}
}
