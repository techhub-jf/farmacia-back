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
	productsPattern = "/products"
)

func (h *Handler) ProductsSetup(router chi.Router) {
	router.Route(productsPattern, func(r chi.Router) {
		r.Get("/", h.GetProducts())
	})
}

func (h *Handler) GetProducts() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		queryStrings := req.URL.Query()

		input := schema.ListProductsRequest{}

		h.getPaginationParams(queryStrings, &input)

		err := input.Validate(schema.ValidateListProductsRequest)
		if err != nil {
			resp := response.BadRequest(err, err.Error())
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		data, err := h.useCase.GetProducts(req.Context(), usecase.GetProductsInput{
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
			TotalItems:   data.TotalProducts,
		}

		products := schema.ConvertProductsToListResponse(data.Products)

		payload := schema.ListProductsOutput{
			Items:    products,
			Metadata: metadata,
		}
		resp := response.OK(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}
