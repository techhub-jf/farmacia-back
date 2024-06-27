package entity

import "time"

type Client struct {
	ID            uint      `json:"id"`
	Reference     string    `json:"reference"`
	FullName      string    `json:"full_name"`
	Birth         time.Time `json:"birth"`
	Cpf           string    `json:"cpf"`
	Rg            string    `json:"rg"`
	Phone         string    `json:"phone"`
	Cep           string    `json:"cep"`
	Address       string    `json:"address"`
	AddressNumber string    `json:"address_number"`
	District      string    `json:"district"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
