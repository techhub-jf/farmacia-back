package entity

import (
	"time"
)

type Delivery struct {
	ID         uint
	Reference  string
	ClientID   uint
	MedicineID uint
	Qty        int32
	TypeID     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
