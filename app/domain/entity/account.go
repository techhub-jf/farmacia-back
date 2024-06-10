package entity

import (
	"time"
)

type Account struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"       validate:"nonzero"`
	Email     string    `json:"email"`
	Secret    string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
