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

// GetClients
// @Summary Get Clients
// @Description Returns Clients
// @Tags Client
// @Produce json
// @Param page query uint64 false "Page number" example(1)
// @Param sort_by query string false "Sort by field" example("name")
// @Param sort_type query string false "Sort type (asc/desc)" example("asc")
// @Param limit query uint64 false "Limit of records per page" example(10)
// @Success 200 {object} []schema.ClientResponse "List of clients"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/farmacia-tech/clients [get]
func (h *Handler) GetClients() http.HandlerFunc {
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
