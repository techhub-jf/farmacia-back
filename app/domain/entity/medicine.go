package entity

import (
	"time"
)

type Medicine struct {
	ID          uint64    `json:"id"`
	Reference   string    `json:"name" validate:"nonzero"`
	Client_id   uint64    `json:"email"`
	Medicine_id uint64    `json:"medicine_id"`
	Qty         int       `json:"qty"`
	Unit_id     int       `json:"unit_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

// type GetMedicinesOutput struct {
// 	Medicines []Medicine `json:"medicines"`
// }
