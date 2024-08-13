package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

func (u *UseCase) UpdateClient(ctx context.Context, clientDTO schema.ClientDTO, id string) (schema.ClientResponse, error) {
	clientID, err := clientDTO.ValidateID(id)
	if err != nil {
		return schema.ClientResponse{}, fmt.Errorf("error updating client: %w", err)
	}

	err = clientDTO.CheckForEmptyFields()
	if err != nil {
		return schema.ClientResponse{}, fmt.Errorf("error updating client: %w", err)
	}

	err = clientDTO.ValidateCpf()
	if err != nil {
		return schema.ClientResponse{}, fmt.Errorf("error updating client: %w", err)
	}

	client := entity.Client{
		ID:            clientID,
		FullName:      clientDTO.FullName,
		Birth:         clientDTO.Birth,
		Cpf:           clientDTO.Cpf,
		Rg:            clientDTO.Rg,
		Phone:         clientDTO.Phone,
		Cep:           clientDTO.Cep,
		Address:       clientDTO.Address,
		AddressNumber: clientDTO.AddressNumber,
		District:      clientDTO.District,
		City:          clientDTO.City,
		State:         clientDTO.State,
		UpdatedAt:     time.Now(),
	}

	outputClient, err := u.ClientsRepository.UpdateClient(ctx, client)
	if err != nil {
		return schema.ClientResponse{}, fmt.Errorf("error updating client: %w", err)
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
