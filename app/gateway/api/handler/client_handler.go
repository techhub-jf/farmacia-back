package handler

import (
	"encoding/json"
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

func (h *Handler) ClientSetup(router chi.Router) {
	router.Route(ClientPattern, func(r chi.Router) {
		r.Get("/", h.ListClients())
		r.Post("/", h.CreateClient())
		r.Put("/{id}", h.UpdateClient())
	})
}

func (h *Handler) CreateClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientDTO schema.ClientDTO

		var resp *response.Response

		err := json.NewDecoder(r.Body).Decode(&clientDTO)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		clientResponse, err := h.useCase.CreateClient(r.Context(), clientDTO)
		if err != nil {
			if errors.Is(err, erring.ErrClientAlreadyExists) {
				resp = response.Conflict(err, err.Error())
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			}

			if errors.Is(err, erring.ErrClientCpfInvalid) {
				resp = response.BadRequest(err, err.Error())
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			}

			if errors.Is(err, erring.ErrClientEmptyFields) {
				resp = response.BadRequest(err, err.Error())
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			}

			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		payload := clientResponse
		resp = response.OK(payload)

		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}

func (h *Handler) UpdateClient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientDTO schema.ClientDTO

		var resp *response.Response

		err := json.NewDecoder(r.Body).Decode(&clientDTO)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		id := chi.URLParam(r, "id")

		clientResponse, err := h.useCase.UpdateClient(r.Context(), clientDTO, id)
		if err != nil {
			switch {
			case errors.Is(err, erring.ErrInvalidID):
				resp = response.BadRequest(err, err.Error())
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return

			case errors.Is(err, erring.ErrResourceNotFound):
				resp = response.NotFound(err, err.Error())
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return

			case errors.Is(err, erring.ErrClientEmptyFields):
				resp = response.BadRequest(err, err.Error())
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return

			case errors.Is(err, erring.ErrClientCpfInvalid):
				resp = response.BadRequest(err, err.Error())
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return

			case errors.Is(err, erring.ErrClientAlreadyExists):
				resp = response.Conflict(err, err.Error())
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return

			default:
				resp = response.InternalServerError(err)
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			}
		}

		payload := clientResponse
		resp = response.OK(payload)

		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
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
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		payload := clients
		resp = response.OK(payload)

		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}
