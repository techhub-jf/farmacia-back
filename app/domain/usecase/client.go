package usecase

import (
	"context"

	"github.com/techhub-jf/farmacia-back/app/domain/erring"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

func (u UseCase) GetClients(ctx context.Context, cqp schema.UnvalidatedClientQueryParams) ([]schema.ClientResponse, error) {
	cqpOut := schema.ValidateParameters(cqp)

	clients, err := u.ClientsRepository.GetClients(ctx, cqpOut)
	if err != nil {
		return []schema.ClientResponse{}, erring.ErrGettingClientsFromDB
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
