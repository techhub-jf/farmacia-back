package usecase

import (
	"context"
)

type ClientOutput struct {
	ID        uint   `json:"id"`
	Reference string `json:"reference"`
	Full_name string `json:"full_name"`
	Cpf       string `json:"cpf"`
	Rg        string `json:"rg"`
	Phone     string `json:"phone"`
}

func (u UseCase) GetClients(ctx context.Context) ([]ClientOutput, error) {
	clients, err := u.ClientsRepository.GetClients(ctx)
	if err != nil {
		// TODO: customize errors
		return []ClientOutput{}, err
	}

	var output []ClientOutput

	for _, client := range clients {
		output = append(output, ClientOutput{
			client.ID,
			client.Reference,
			client.Full_name,
			client.Cpf,
			client.Rg,
			client.Phone,
		})
	}

	return output, nil
}
