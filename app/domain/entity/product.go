package entity

import (
	"time"
)

type Product struct {
	ID              uint      `json:"id"`
	Reference       string    `json:"name" validate:"nonzero"`
	Qty             uint      `json:"qty"`
	Description     string    `json:"description"`
	ActivePrinciple string    `json:"active_principle"`
	UnitID          uint      `json:"unit_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}
