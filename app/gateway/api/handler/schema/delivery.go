package schema

import (
	"errors"
	"time"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

type ListDeliveriesRequest = dto.Pagination

var validSearchFields = map[string]bool{
	"reference":  true,
	"qty":        true,
	"id":         true,
	"created_at": true,
}

func ValidateListDeliveriesRequest(input ListDeliveriesRequest) error {
	if input.Page < 1 {
		return errors.New("page must be greater than 0")
	}

	if input.ItemsPerPage < 1 || input.ItemsPerPage > 100 {
		return errors.New("itemsPerPage must be between 1 and 100")
	}

	if _, found := validSearchFields[input.SortBy]; !found {
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
	CreatedAt time.Time `json:"created_at"`
}

type ListDeliveriesOutput = PaginatedResponse[ListDeliveriesResponse]

func ConvertDeliveriesToListResponse(deliveries []entity.Delivery) []ListDeliveriesResponse {
	parsedDeliveries := []ListDeliveriesResponse{}

	for _, delivery := range deliveries {
		parsedDeliveries = append(parsedDeliveries, ListDeliveriesResponse{
			ID:        delivery.ID,
			Reference: delivery.Reference,
			Qty:       delivery.Qty,
			CreatedAt: delivery.CreatedAt,
		})
	}
	return parsedDeliveries
}
