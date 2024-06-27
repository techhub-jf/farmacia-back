package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
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

		page := values.Get("page")
		sortBy := values.Get("sort_by")
		sortType := values.Get("sort_type")
		limit := values.Get("limit")

		cqp := usecase.ClientQueryParametersInput{
			Page:     page,
			SortBy:   sortBy,
			SortType: sortType,
			Limit:    limit,
		}

		clients, err := h.useCase.GetClients(r.Context(), cqp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		payload := clients
		resp := response.OK(payload)

		err = rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	}
}
