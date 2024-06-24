package entity

import (
	"time"
)

type Delivery struct {
	ID         uint      `json:"id"`
	Reference  string    `json:"reference"`
	ClientId   uint      `json:"-"`
	MedicineId uint      `json:"-"`
	Qty        int32     `json:"qty"`
	UnitId     uint      `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"-"`
	DeletedAt  time.Time `json:"-"`
}
