package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/techhub-jf/farmacia-back/app/domain/erring"
	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest/response"
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
		creds := &schema.LoginRequest{}
		err := json.NewDecoder(req.Body).Decode(creds)
		var resp *response.Response
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}
		input := usecase.LoginInput{
			Email:        creds.Email,
			Password:     creds.Password,
			JwtSecretKey: h.cfg.JwtSecretKey,
		}
		output, err := h.useCase.Login(req.Context(), input)

		if err != nil {
			switch err {
			case erring.ErrLoginUserNotFound:
				resp = response.NotFound(err, "401", err.Error())
				rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
				return
			case erring.ErrLoginUnauthorized:
				resp = response.Unauthorized(err.Error())
				rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
				return
			case erring.ErrLoginTokenNotCreated:
				resp = response.InternalServerError(fmt.Errorf("internal error"))
				rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
				return
			}
		}

		respBody := &schema.LoginResponse{
			Token: output.Token,
			Email: output.Account.Email,
			Name:  output.Account.Name,
		}

		if err != nil {
			resp = response.InternalServerError(fmt.Errorf("internal error"))
			rest.SendJSON(rw, resp.Status, err.Error(), resp.Headers)
			return
		}
		resp = response.OK(respBody)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
	}

}
