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
		r.Get("/reference/{reference}", h.GetDeliveryByReference())
	})
}

// GetDeliveries
// @Summary Get Deliveries
// @Description Returns deliveries
// @Tags Delivery
// @Produce json
// @Param page query int false "Page number" example(1)
// @Param items_per_page query int false "Number of items per page" example(10)
// @Param sort_by query string false "Field to sort by" example("name")
// @Param sort_type query string false "Type of sorting (asc/desc)" example("asc")
// @Success 200 {object} schema.ListDeliveriesOutput "List of deliveries"s
// @Failure 500 "Internal Server Error"
// @Router /api/v1/farmacia-tech/deliveries [get]
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

// CreateDelivery creates a new delivery
// @Summary Create Delivery
// @Description Create a new delivery record
// @Tags Delivery
// @Accept json
// @Produce json
// @Param delivery body schema.CreateDeliveryRequest true "Delivery data"
// @Success 201 {object} schema.CreateDeliveryResponse "Created delivery"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/farmacia-tech/deliveries [post]
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

// GetDeliveryByReference retrieves a delivery by its reference
// @Summary Get Delivery by Reference
// @Description Get details of a specific delivery using its reference
// @Tags Delivery
// @Produce json
// @Param reference path string true "Delivery Reference"
// @Success 200 {object} schema.GetDeliveryByReferenceResponse "Delivery details"
// @Failure 404 "Not Found"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/farmacia-tech/deliveries/reference/{reference} [get]
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
