package schema

import (
	"time"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
)

type ListDeliveriesInput = dto.Pagination

type ListDeliveriesResponse struct {
	ID        uint      `json:"id"`
	Reference string    `json:"reference"`
	Qty       int32     `json:"qty"`
	UnitID    uint      `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type ListDeliveriesOutput = PaginationResponse[*ListDeliveriesResponse]
