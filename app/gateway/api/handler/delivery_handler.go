package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

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
		r.Delete("/{id}", h.DeleteDelivery())
		r.Get("/reference/{reference}", h.GetDeliveryByReference())
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
		useCaseInput.Delivery.MedicineID = deliveryBody.MedicineID
		useCaseInput.Delivery.UnitID = deliveryBody.UnitID
		useCaseInput.Delivery.ClientID = deliveryBody.ClientID

		var reference string

		for {
			reference = strconv.Itoa(rand.Intn(schema.MaxReference) + schema.MinReference) //nolint:gosec

			_, err = h.useCase.GetDeliveryByReference(req.Context(), usecase.GetDeliveryByReferenceInput{
				Reference: reference,
			})
			if err != nil {
				if strings.Contains(err.Error(), "no rows in result set") {
					useCaseInput.Delivery.Reference = reference

					break
				}

				resp := response.InternalServerError(err)
				rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			}
		}

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

func (h *Handler) GetDeliveryByReference() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		reference := chi.URLParam(req, "reference")

		data, err := h.useCase.GetDeliveryByReference(req.Context(), usecase.GetDeliveryByReferenceInput{
			Reference: reference,
		})
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				resp := response.NotFound(err, err.Error())
				rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			}

			resp := response.InternalServerError(err)

			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		payload := schema.GetDeliveryByReferenceResponse{}
		delivery := schema.ConvertDeliveryToGetResponse(data.Delivery)
		payload.Delivery = delivery

		resp := response.OK(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}

func (h *Handler) DeleteDelivery() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			resp := response.BadRequest(err, "Invalid ID")
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		err = h.useCase.DeleteDelivery(req.Context(), usecase.DeleteDeliveryInput{
			Id: int32(idInt),
		})
		if err != nil {
			if notFound := strings.Contains(err.Error(), "no rows in result set") || strings.Contains(err.Error(), "delivery already deleted"); notFound {
				resp := response.NotFound(err, "Delivery not found")
				rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			}

			resp := response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		resp := response.NoContent()
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}
