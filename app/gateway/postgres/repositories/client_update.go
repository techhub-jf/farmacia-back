package repositories

import (
	"context"
	"fmt"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/domain/erring"
)

const (
	clientCPFAlreadyRegisteredClause = `
	SELECT EXISTS (SELECT 1 FROM client WHERE cpf = $1 AND id <> $2 AND deleted_at IS NULL)
	`

	clientIdExistsClause = `
	SELECT EXISTS (SELECT 1 FROM client WHERE id = $1 AND deleted_at IS NULL)
	`

	updateClientClause = `
	UPDATE client
	SET full_name = $1, birth = $2, cpf = $3, rg = $4, phone = $5, cep = $6, address = $7, address_number = $8, district = $9, city = $10, state = $11
	WHERE id = $12 AND deleted_at IS NULL
	RETURNING id, reference, full_name, cpf, rg, phone
`
)

func (r *ClientsRepository) UpdateClient(ctx context.Context, client entity.Client) (entity.Client, error) {
	const operation = "Repository.ClientsRepository.UpdateClient"

	var clientExists bool

	err := r.Pool.QueryRow(
		ctx,
		clientIdExistsClause,
		client.ID,
	).Scan(
		&clientExists,
	)
	if err != nil {
		return entity.Client{}, fmt.Errorf("%s -> %w", operation, err)
	}

	if !clientExists {
		return entity.Client{}, fmt.Errorf("%s -> %w", operation, erring.ErrResourceNotFound)
	}

	var clientCPFAlreadyRegistered bool
	err = r.Pool.QueryRow(
		ctx,
		clientCPFAlreadyRegisteredClause,
		client.Cpf,
		client.ID,
	).Scan(
		&clientCPFAlreadyRegistered,
	)
	if err != nil {
		return entity.Client{}, fmt.Errorf("%s -> %w", operation, err)
	}

	if clientCPFAlreadyRegistered {
		return entity.Client{}, fmt.Errorf("%s -> %w", operation, erring.ErrClientAlreadyExists)
	}

	args := []interface{}{
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
		client.ID,
	}

	var clientResponse entity.Client

	err = r.Pool.QueryRow(
		ctx,
		updateClientClause,
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
