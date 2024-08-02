package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest/response"
)

const (
	medicinePattern = "/medicine"
)

// GetMedicines retrieves a list of medicines
// @Summary Retrieve Medicines
// @Description Get a list of all medicines available
// @Tags medicines
// @Accept  json
// @Produce json
// @Success 200 {array} Medicine
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/farmacia-tech/medicines [get]
func (h *Handler) GetMedicinesSetup(router chi.Router) {
	router.Route(medicinePattern, func(r chi.Router) {
		r.Get("/", h.GetMedicines())
	})
}

func (h *Handler) GetMedicines() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		medicines, err := h.useCase.MedicinesRepository.GetMedicines(req.Context())
		if err != nil {
			resp := response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
			return
		}

		payload := medicines
		resp := response.OK(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers)
	}
}
