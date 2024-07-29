package repositories

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/domain/erring"
)

const (
	clientCPFExistsClause = `
	SELECT EXISTS (SELECT 1 FROM client WHERE cpf = $1 AND deleted_at IS NULL)
	`

	createClientClause = `
	INSERT INTO client (reference, full_name, birth, cpf, rg, phone, cep, address, address_number, district, city, state)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	RETURNING id, reference, full_name, cpf, rg, phone
	`
)

func (r *ClientsRepository) CreateClient(ctx context.Context, client entity.Client) (entity.Client, error) {
	const operation = "Repository.ClientsRepository.CreateClient"

	var clientAlreadyExists bool

	err := r.Pool.QueryRow(
		ctx,
		clientCPFExistsClause,
		client.Cpf,
	).Scan(
		&clientAlreadyExists,
	)
	if err != nil {
		return entity.Client{}, fmt.Errorf("%s -> %w", operation, err)
	}

	if clientAlreadyExists {
		return entity.Client{}, fmt.Errorf("%s -> %w", operation, erring.ErrClientAlreadyExists)
	}

	args := []interface{}{
		client.Reference,
		client.FullName,
		client.Birth,
		client.Cpf,
		client.Rg,
		client.Phone,
		client.Cep,
		client.Address,
		client.AddressNumber,
		client.District,
		client.City,
		client.State,
	}

	var clientResponse entity.Client

	err = r.Pool.QueryRow(
		ctx,
		createClientClause,
		args...,
	).Scan(
		&clientResponse.ID,
		&clientResponse.Reference,
		&clientResponse.FullName,
		&clientResponse.Cpf,
		&clientResponse.Rg,
		&clientResponse.Phone,
	)
	if err != nil {
		return entity.Client{}, err
	}

	return clientResponse, nil
}
