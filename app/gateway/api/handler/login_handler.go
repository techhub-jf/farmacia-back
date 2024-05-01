package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest/response"
	"golang.org/x/crypto/bcrypt"
)

const (
	loginPattern = "/auth"
)

func (h *Handler) LoginSetup(router chi.Router) {
	router.Route(loginPattern, func(r chi.Router) {
		r.Post("/", h.login())
	})
}

func (h *Handler) login() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		creds := &dto.LoginRequest{}
		err := json.NewDecoder(req.Body).Decode(creds)
		var resp *response.Response
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}

		account, err := h.useCase.AccountsRepository.GetAccountByEmail(req.Context(), creds.Email)
		if err != nil {
			resp = response.NotFound(err, "401", "user not found")
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(creds.Password)); err != nil {
			resp = response.Unauthorized()
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}
		tokenString, err := createToken(account.ID, h.cfg.JwtSecretKey)
		if err != nil {
			resp = response.InternalServerError(fmt.Errorf("internal error"))
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}

		payload := simplejson.New()
		payload.Set("token", tokenString)
		resp = response.OK(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
	}

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
