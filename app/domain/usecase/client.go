package usecase

import (
	"context"
	"fmt"
	"math/rand/v2"
	"strconv"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/domain/erring"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

func (u *UseCase) CreateClient(ctx context.Context, clientDTO schema.ClientDTO) (schema.ClientResponse, error) {
	err := clientDTO.CheckForEmptyFields()
	if err != nil {
		return schema.ClientResponse{}, erring.ErrClientEmptyFields
	}

	err = clientDTO.ValidateCpf()
	if err != nil {
		return schema.ClientResponse{}, erring.ErrClientCpfInvalid
	}

	client := entity.Client{
		Reference:     fmt.Sprint(rand.Int32N(900000) + 100000),
		FullName:      clientDTO.FullName,
		Cpf:           clientDTO.Cpf,
		Rg:            clientDTO.Rg,
		Phone:         clientDTO.Phone,
		Cep:           clientDTO.Cep,
		Address:       clientDTO.Address,
		AddressNumber: clientDTO.AddressNumber,
		District:      clientDTO.District,
		City:          clientDTO.City,
		State:         clientDTO.State,
	}

	outputClient, err := u.ClientsRepository.CreateClient(ctx, client)
	if err != nil {
		return schema.ClientResponse{}, err
	}

	return schema.ClientResponse{
		ID:        outputClient.ID,
		Reference: outputClient.Reference,
		FullName:  outputClient.FullName,
		Cpf:       outputClient.Cpf,
		Rg:        outputClient.Rg,
		Phone:     outputClient.Phone,
	}, nil
}

func (u *UseCase) UpdateClient(ctx context.Context, clientDTO schema.ClientDTO, id string) (schema.ClientResponse, error) {
	clientID, err := strconv.ParseUint(id, 10, 0)
	if err != nil {
		return schema.ClientResponse{}, err
	}

	err = clientDTO.CheckForEmptyFields()
	if err != nil {
		return schema.ClientResponse{}, erring.ErrClientEmptyFields
	}

	err = clientDTO.ValidateCpf()
	if err != nil {
		return schema.ClientResponse{}, erring.ErrClientCpfInvalid
	}

	client := entity.Client{
		ID:            uint(clientID),
		FullName:      clientDTO.FullName,
		Cpf:           clientDTO.Cpf,
		Rg:            clientDTO.Rg,
		Phone:         clientDTO.Phone,
		Cep:           clientDTO.Cep,
		Address:       clientDTO.Address,
		AddressNumber: clientDTO.AddressNumber,
		District:      clientDTO.District,
		City:          clientDTO.City,
		State:         clientDTO.State,
	}

	outputClient, err := u.ClientsRepository.UpdateClient(ctx, client)
	if err != nil {
		return schema.ClientResponse{}, err
	}

	return schema.ClientResponse{
		ID:        outputClient.ID,
		Reference: outputClient.Reference,
		FullName:  outputClient.FullName,
		Cpf:       outputClient.Cpf,
		Rg:        outputClient.Rg,
		Phone:     outputClient.Phone,
	}, nil
}

func (u *UseCase) GetClients(ctx context.Context, cqp schema.ClientQueryParams) ([]schema.ClientResponse, error) {
	clients, err := u.ClientsRepository.GetClients(ctx, cqp)
	if err != nil {
		return []schema.ClientResponse{}, err
	}

	clientListOutput := make([]schema.ClientResponse, 0)

	for _, client := range clients {
		clientListOutput = append(clientListOutput, schema.ClientResponse{
			ID:        client.ID,
			Reference: client.Reference,
			FullName:  client.FullName,
			Cpf:       client.Cpf,
			Rg:        client.Rg,
			Phone:     client.Phone,
		})
	}

	return clientListOutput, nil
}
