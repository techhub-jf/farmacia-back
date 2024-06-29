package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest/response"
)

const (
	deliveryPattern = "/deliveries"
)

func (h *Handler) ListDeliveriesSetup(router chi.Router) {
	router.Route(deliveryPattern, func(r chi.Router) {
		r.Get("/", h.ListDeliveries())
	})
}

func (h *Handler) ListDeliveries() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		queryStrings := req.URL.Query()

		input := schema.ListDeliveriesInput{}

		h.getPaginationParams(queryStrings, &input)

		deliveries, err := h.useCase.GetDeliveries(req.Context(), input)
		if err != nil {
			resp := response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		payload := deliveries
		resp := response.OK(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}
