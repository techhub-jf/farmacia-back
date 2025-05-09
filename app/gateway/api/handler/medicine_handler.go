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

func (h *Handler) GetMedicinesSetup(router chi.Router) {
	router.Route(medicinePattern, func(r chi.Router) {
		r.Get("/", h.GetMedicines())
	})
}

// GetMedicines
// @Summary Get Medicines
// @Description Returns medicines
// @Tags Medicine
// @Produce json
// @Success 200 {object} []entity.Medicine "List of medicines"
// @Failure 500 "Internal Server Error"
// @Router /api/v1/farmacia-tech/medicine [get]
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
