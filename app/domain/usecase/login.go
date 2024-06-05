package usecase

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/techhub-jf/farmacia-back/app/domain/entity"
	"github.com/techhub-jf/farmacia-back/app/domain/erring"
	"golang.org/x/crypto/bcrypt"
)

type LoginOutput struct {
	Account entity.Account
	Token   string
}

type LoginInput struct {
	Email        string
	Password     string
	JwtSecretKey string
}

func (u *UseCase) Login(ctx context.Context, input LoginInput) (LoginOutput, error) {
	account, err := u.AccountsRepository.GetAccountByEmail(ctx, input.Email)
	if err != nil {
		return LoginOutput{}, erring.ErrLoginUserNotFound
	}

	if err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(input.Password)); err != nil {
		return LoginOutput{}, erring.ErrLoginUnauthorized
	}

	tokenString, err := createToken(account.ID, input.JwtSecretKey)
	if err != nil {
		return LoginOutput{}, erring.ErrLoginTokenNotCreated
	}

	return LoginOutput{
		Account: account,
		Token:   tokenString,
	}, nil
}

func createToken(user uint, jwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": user,
			"exp":  time.Now().Add(time.Hour * 24).Unix(),
		})

	jwtKey := []byte(jwtSecret)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
