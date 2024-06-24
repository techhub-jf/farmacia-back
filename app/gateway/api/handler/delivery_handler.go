package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest/response"
)

const (
	deliveryPattern = "/deliveries"
)

func (h *Handler) GetDeliveriesSetup(router chi.Router) {
	router.Route(deliveryPattern, func(r chi.Router) {
		r.Get("/", h.GetAllDeliveries())
	})
}

func (h *Handler) GetAllDeliveries() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		input := usecase.GetDeliveriesInput{
			Page:     1,
			SortBy:   "reference",
			SortType: "ASC",
		}

		deliveries, err := h.useCase.GetDeliveries(req.Context(), input)
		if err != nil {
			print("error:", err.Error())
			resp := response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}

		payload := deliveries
		resp := response.OK(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
	}
}
