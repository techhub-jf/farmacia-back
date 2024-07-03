package handler

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
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

		input := schema.ListDeliveriesRequest{}

		h.getPaginationParams(queryStrings, &input)

		err := input.Validate(schema.ValidateListDeliveriesRequest)
		if err != nil {
			resp := response.BadRequest(err, err.Error())
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		data, err := h.useCase.GetDeliveries(req.Context(), usecase.GetDeliveriesInput{
			Pagination: input,
		})
		if err != nil {
			resp := response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		metadata := schema.Meta{
			ItemsPerPage: input.ItemsPerPage,
			CurrentPage:  input.Page,
			TotalItems:   data.TotalDeliveries,
		}

		deliveries := schema.ConvertDeliveriesToListResponse(data.Deliveries)

		payload := schema.ListDeliveriesOutput{
			Items:    deliveries,
			Metadata: metadata,
		}
		resp := response.OK(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}
