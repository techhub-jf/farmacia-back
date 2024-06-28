package entity

import (
	"time"
)

type Delivery struct {
	ID         uint
	Reference  string
	ClientId   uint
	MedicineId uint
	Qty        int32
	UnitId     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
