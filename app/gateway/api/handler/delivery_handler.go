package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest/response"
)

const (
	deliveryPattern = "/deliveries"
)

func (h *Handler) DeliveriesSetup(router chi.Router) {
	router.Route(deliveryPattern, func(r chi.Router) {
		r.Get("/", h.ListDeliveries())
		r.Post("/", h.CreateDelivery())
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

func (h *Handler) CreateDelivery() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		deliveryBody := &schema.CreateDeliveryRequest{}
		err := json.NewDecoder(req.Body).Decode(deliveryBody)

		var resp *response.Response

		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		err = schema.ValidateCreateDeliveryRequest(deliveryBody)
		if err != nil {
			resp := response.BadRequest(err, err.Error())
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		useCaseInput := usecase.CreateDeliveryInput{}
		useCaseInput.Delivery.Qty = deliveryBody.Qty
		useCaseInput.Delivery.Reference = strconv.Itoa(rand.Intn(999999) + 100000)
		useCaseInput.Delivery.MedicineID = deliveryBody.MedicineID
		useCaseInput.Delivery.UnitID = deliveryBody.UnitID
		useCaseInput.Delivery.ClientID = deliveryBody.ClientID

		data, err := h.useCase.CreateDelivery(req.Context(), useCaseInput)
		if err != nil {
			resp := response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		payload := schema.CreateDeliveryResponse{}
		payload.Delivery = schema.ConvertDeliveryToCreateResponse(data.Delivery)

		resp = response.Created(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}
