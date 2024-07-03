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
	ClientPattern = "/client"
)

func (h *Handler) ClientRoutes(router chi.Router) {
	router.Route(ClientPattern, func(r chi.Router) {
		r.Get("/list", h.ListClients())
		r.Post("/create", h.CreateClient())
		r.Put("/update/{id}", h.UpdateClient())
	})
}

func (h *Handler) CreateClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("create client")
	}
}

func (h *Handler) UpdateClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("update client")
	}
}

func (h *Handler) ListClients() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()

		page := values.Get("page")
		sortBy := values.Get("sort_by")
		sortType := values.Get("sort_type")
		limit := values.Get("limit")

		var cqp schema.ClientQueryParams

		cqp.ValidateParameters(page, sortBy, sortType, limit)

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
