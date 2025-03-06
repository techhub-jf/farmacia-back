package usecase

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

func (u *UseCase) CreateClient(ctx context.Context, clientDTO schema.ClientDTO) (schema.ClientResponse, error) {
	err := clientDTO.CheckForEmptyFields()
	if err != nil {
		return schema.ClientResponse{}, fmt.Errorf("error creating client: %w", err)
	}

	err = clientDTO.ValidateCpf()
	if err != nil {
		return schema.ClientResponse{}, fmt.Errorf("error creating client: %w", err)
	}

	reference, err := generateReferenceNumber()
	if err != nil {
		return schema.ClientResponse{}, fmt.Errorf("error creating client: %w", err)
	}

	client := entity.Client{
		Reference:     reference,
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
	}

	outputClient, err := u.ClientsRepository.CreateClient(ctx, client)
	if err != nil {
		return schema.ClientResponse{}, fmt.Errorf("error creating client: %w", err)
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

func generateReferenceNumber() (string, error) {
	const (
		minReferenceSize = 100_000
		maxReferenceSize = 900_000
	)

	newBigInt := big.NewInt(maxReferenceSize)

	randomNumber, err := rand.Int(rand.Reader, newBigInt)
	if err != nil {
		return "", fmt.Errorf("failed to generate reference number: %w", err)
	}

	return strconv.FormatUint(randomNumber.Uint64()+minReferenceSize, 10), nil
}
