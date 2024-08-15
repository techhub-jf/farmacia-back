package entity

import (
	"time"
)

type Type struct {
	ID        uint      `json:"id"`
	Reference string    `json:"reference"   validate:"nonzero"`
	Label     string    `json:"label"   `
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
