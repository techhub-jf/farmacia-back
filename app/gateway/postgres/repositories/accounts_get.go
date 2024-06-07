package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
)

const getAccountByEmailClause = `
SELECT
name,
email,
secret
FROM account
WHERE email = $1
AND deleted_at IS NULL
`

func (r *AccountsRepository) GetAccountByEmail(ctx context.Context, email string) (entity.Account, error) {
	const (
		operation = "Repository.AccountsRepository.GetAccountByEmail"
	)

	var account entity.Account

	err := r.Client.Pool.QueryRow(
		ctx,
		getAccountByEmailClause,
		email,
	).Scan(
		&account.Name,
		&account.Email,
		&account.Secret,
	)
	if err != nil {
		if errors.Is(pgx.ErrNoRows, err) {
			return entity.Account{}, fmt.Errorf("%s -> %w", operation, err)
		}

		return entity.Account{}, fmt.Errorf("%s -> %w", operation, err)
	}

	return account, nil
}
