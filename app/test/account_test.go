package test

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	mock_usecase "github.com/techhub-jf/farmacia-back/app/test/mock"
)

func TestGetAccountByEmail(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_usecase.NewMockaccountsRepository(ctrl)

	email := "teste@email.com"

	mock.EXPECT().GetAccountByEmail(ctx, email).Return(entity.Account{
		Email:  "teste@email.com",
		Secret: "fbewjfkcbnldksn",
	}, nil)

	account, err := mock.GetAccountByEmail(ctx, email)
	if err != nil {
		t.FailNow()

		return
	}

	if account.Email != email {
		t.FailNow()

		return
	}
}

func TestGetAccountByEmailFail(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_usecase.NewMockaccountsRepository(ctrl)

	email := "teste@email.com"

	mock.EXPECT().GetAccountByEmail(ctx, email).Return(entity.Account{}, errors.New("user not found"))

	account, err := mock.GetAccountByEmail(ctx, email)

	if err == nil {
		t.FailNow()

		return
	}

	if account.Email == email {
		t.FailNow()

		return
	}
}
