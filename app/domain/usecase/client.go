package usecase

import (
	"context"
	"strconv"
	"strings"

	"github.com/techhub-jf/farmacia-back/app/domain/erring"
)

type ClientOutput struct {
	ID        uint   `json:"id"`
	Reference string `json:"reference"`
	FullName  string `json:"full_name"`
	Cpf       string `json:"cpf"`
	Rg        string `json:"rg"`
	Phone     string `json:"phone"`
}

type ClientQueryParametersInput struct {
	Page     string
	SortBy   string
	SortType string
	Limit    string
}

type ClientQueryParametersOutput struct {
	Page     uint64
	SortBy   string
	SortType string
	Limit    uint64
}

func (u UseCase) GetClients(ctx context.Context, cqp ClientQueryParametersInput) ([]ClientOutput, error) {
	cqpOut := validateParameters(cqp)

	clients, err := u.ClientsRepository.GetClients(ctx, cqpOut)
	if err != nil {
		return []ClientOutput{}, erring.ErrGettingClientsFromDB
	}

	clientListOutput := make([]ClientOutput, 0)

	for _, client := range clients {
		clientListOutput = append(clientListOutput, ClientOutput{
			client.ID,
			client.Reference,
			client.FullName,
			client.Cpf,
			client.Rg,
			client.Phone,
		})
	}

	return clientListOutput, nil
}

func validateParameters(cqp ClientQueryParametersInput) ClientQueryParametersOutput {
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

	var outputSortType string

	cqp.SortType = strings.ToUpper(cqp.SortType)

	if cqp.SortType == "DESC" {
		outputSortType = cqp.SortType
	} else {
		outputSortType = "ASC"
	}

	outputLimit, err := strconv.ParseUint(cqp.Limit, 10, 32)
	if err != nil {
		outputLimit = 10
	}

	return ClientQueryParametersOutput{
		Page:     outputPage,
		SortBy:   outputSortBy,
		SortType: outputSortType,
		Limit:    outputLimit,
	}
}
