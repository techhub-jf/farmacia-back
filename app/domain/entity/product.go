package entity

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id"`
	Reference   string    `json:"reference"   validate:"nonzero"`
	Stock       uint      `json:"stock"`
	Description string    `json:"description"`
	Brand       string    `json:"brand"`
	TypeID      uint      `json:"type_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
