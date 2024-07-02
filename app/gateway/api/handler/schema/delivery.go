package schema

import (
	"errors"
	"time"
)

type ListDeliveriesInput = Pagination

var validSearchFields = map[string]bool{
	"reference":  true,
	"qty":        true,
	"id":         true,
	"created_at": true,
}

func ValidateListDeliveriesInput(input ListDeliveriesInput) error {
	if input.Page < 1 {
		return errors.New("page must be greater than 0")
	}

	if input.ItemsPerPage < 1 || input.ItemsPerPage > 100 {
		return errors.New("itemsPerPage must be between 1 and 100")
	}

	_, found := validSearchFields[input.SortBy]
	if !found {
		return errors.New("invalid sortBy param")
	}

	if input.SortType != "ASC" && input.SortBy != "DESC" {
		return errors.New("sortType param must be 'ASC' or 'DESC'")
	}

	return nil
}

type ListDeliveriesResponse struct {
	ID        uint      `json:"id"`
	Reference string    `json:"reference"`
	Qty       int32     `json:"qty"`
	UnitID    uint      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type ListDeliveriesOutput = PaginatedResponse[ListDeliveriesResponse]
