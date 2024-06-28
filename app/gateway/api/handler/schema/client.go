package schema

import (
	"strconv"
	"strings"
)

type ClientResponse struct {
	ID        uint   `json:"id"`
	Reference string `json:"reference"`
	FullName  string `json:"full_name"`
	Cpf       string `json:"cpf"`
	Rg        string `json:"rg"`
	Phone     string `json:"phone"`
}

type UnvalidatedClientQueryParams struct {
	Page     string
	SortBy   string
	SortType string
	Limit    string
}

type ValidatedClientQueryParams struct {
	Page     uint64
	SortBy   string
	SortType string
	Limit    uint64
}

func ValidateParameters(cqp UnvalidatedClientQueryParams) ValidatedClientQueryParams {
	outputPage, err := strconv.ParseUint(cqp.Page, 10, 32)
	if err != nil {
		outputPage = 1
	}

	var outputSortBy string
	if cqp.SortBy == "reference" ||
		cqp.SortBy == "full_name" ||
		cqp.SortBy == "cpf" ||
		cqp.SortBy == "rg" ||
		cqp.SortBy == "phone" {

		outputSortBy = cqp.SortBy
	} else {
		outputSortBy = "id"
	}

	cqp.SortType = strings.ToUpper(cqp.SortType)

	var outputSortType string
	if cqp.SortType == "DESC" {
		outputSortType = cqp.SortType
	} else {
		outputSortType = "ASC"
	}

	outputLimit, err := strconv.ParseUint(cqp.Limit, 10, 32)
	if err != nil {
		outputLimit = 10
	}

	return ValidatedClientQueryParams{
		Page:     outputPage,
		SortBy:   outputSortBy,
		SortType: outputSortType,
		Limit:    outputLimit,
	}
}
