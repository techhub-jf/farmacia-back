package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	mock_usecase "github.com/techhub-jf/farmacia-back/app/test/mock"
	"go.uber.org/mock/gomock"
)

func TestGetAccountByEmail(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_usecase.NewMockaccountsRepository(ctrl)

	email := "teste@email.com"

	m.EXPECT().GetAccountByEmail(ctx, email).Return(entity.Account{
		Email:  "teste@email.com",
		Secret: "fbewjfkcbnldksn",
	}, nil)

	account, err := m.GetAccountByEmail(ctx, email)

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
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_usecase.NewMockaccountsRepository(ctrl)

	email := "teste@email.com"

	m.EXPECT().GetAccountByEmail(ctx, email).Return(entity.Account{}, fmt.Errorf("user not found"))

	account, err := m.GetAccountByEmail(ctx, email)

	if err == nil {
		t.FailNow()
		return
	}

	if account.Email == email {
		t.FailNow()
		return
	}

}
