package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/techhub-jf/farmacia-back/app/domain/erring"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest/response"
)

const (
	clientPattern = "/clients"
)

func (h *Handler) ListClients(router chi.Router) {
	router.Route(clientPattern, func(r chi.Router) {
		r.Get("/", h.GetClients())
	})
}

func (h *Handler) GetClients() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()

		cqp := schema.UnvalidatedClientQueryParams{
			Page:     values.Get("page"),
			SortBy:   values.Get("sort_by"),
			SortType: values.Get("sort_type"),
			Limit:    values.Get("limit"),
		}

		var resp *response.Response

		clients, err := h.useCase.GetClients(r.Context(), cqp)
		if err != nil && errors.Is(err, erring.ErrGettingClientsFromDB) {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		payload := clients
		resp = response.OK(payload)

		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}
