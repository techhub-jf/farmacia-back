package schema

import (
	"strconv"
	"strings"
	"time"
)

type CreateClientRequest struct {
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
}

type ClientResponse struct {
	ID        uint   `json:"id"`
	Reference string `json:"reference"`
	FullName  string `json:"full_name"`
	Cpf       string `json:"cpf"`
	Rg        string `json:"rg"`
	Phone     string `json:"phone"`
}

type ClientQueryParams struct {
	Page     uint64
	SortBy   string
	SortType string
	Limit    uint64
}

func (cqp *ClientQueryParams) ValidateParameters(page string, sortBy string, sortType string, limit string) {
	outputPage, err := strconv.ParseUint(page, 10, 32)
	if err != nil {
		outputPage = 1
	}

	var outputSortBy string
	if sortBy == "reference" ||
		sortBy == "full_name" ||
		sortBy == "cpf" ||
		sortBy == "rg" ||
		sortBy == "phone" {
		outputSortBy = sortBy
	} else {
		outputSortBy = "id"
	}

	sortType = strings.ToUpper(sortType)

	var outputSortType string
	if sortType == "DESC" {
		outputSortType = sortType
	} else {
		outputSortType = "ASC"
	}

	outputLimit, err := strconv.ParseUint(limit, 10, 32)
	if err != nil {
		outputLimit = 10
	}

	cqp.Page = outputPage
	cqp.SortBy = outputSortBy
	cqp.SortType = outputSortType
	cqp.Limit = outputLimit
}
