package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest/response"
)

const (
	clientPattern = "/clients"
)

func (h *Handler) GetClientsSetup(router chi.Router) {
	router.Route(clientPattern, func (r chi.Router){
		r.Get("/", h.GetClients())
	})
}

func (h *Handler) GetClients() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clients, err := h.useCase.GetClients(r.Context())

		if (err != nil) {
			resp := response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		payload := clients
		resp := response.OK(payload)
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}

