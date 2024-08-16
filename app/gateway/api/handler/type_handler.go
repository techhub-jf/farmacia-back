package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"

	"github.com/techhub-jf/farmacia-back/app/domain/erring"
	"github.com/techhub-jf/farmacia-back/app/domain/usecase"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/rest/response"
)

const (
	typesPattern = "/types"
)

func (h *Handler) TypesSetup(router chi.Router) {
	router.Route(typesPattern, func(r chi.Router) {
		r.Get("/", h.GetTypes())
		r.Post("/", h.CreateType())
		r.Put("/{id}", h.UpdateType())
		r.Delete("/{id}", h.DeleteType())
	})
}

func (h *Handler) GetTypes() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		queryStrings := req.URL.Query()

		input := schema.ListTypesRequest{}

		h.getPaginationParams(queryStrings, &input.Pagination)
		input.Search = h.readString(queryStrings, "search", "")

		err := input.Pagination.Validate(schema.ValidateListTypesRequest)
		if err != nil {
			resp := response.BadRequest(err, err.Error())
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		data, err := h.useCase.GetTypes(req.Context(), usecase.GetTypesInput{
			Pagination: input.Pagination,
			Search:     input.Search,
		})
		if err != nil {
			resp := response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		metadata := schema.Meta{
			ItemsPerPage: input.Pagination.ItemsPerPage,
			CurrentPage:  input.Pagination.Page,
			TotalItems:   data.TotalTypes,
		}

		types := schema.ConvertTypesToListResponse(data.Types)

		payload := schema.ListTypesOutput{
			Items:    types,
			Metadata: metadata,
		}
		resp := response.OK(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}

func (h *Handler) CreateType() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		typeBody := &schema.CreateTypeRequest{}
		err := json.NewDecoder(req.Body).Decode(typeBody)

		var resp *response.Response

		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		useCaseInput := usecase.CreateTypeInput{}
		useCaseInput.Type.Label = typeBody.Label

		_, err = h.useCase.GetTypeByLabel(req.Context(), usecase.GetTypeByLabelInput{
			Label: useCaseInput.Type.Label,
		})

		if err == nil {
			resp := response.BadRequest(erring.ErrLabelExists, erring.ErrLabelExists.Message)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		var reference string

		for {
			reference = strconv.Itoa(rand.Intn(schema.MaxReference) + schema.MinReference) //nolint:gosec

			_, err = h.useCase.GetTypeByReference(req.Context(), usecase.GetTypeByReferenceInput{
				Reference: reference,
			})
			if err != nil {
				if strings.Contains(err.Error(), "no rows in result set") {
					useCaseInput.Type.Reference = reference

					break
				}

				resp := response.InternalServerError(err)
				rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

				return
			}
		}

		data, err := h.useCase.CreateType(req.Context(), useCaseInput)
		if err != nil {
			resp := response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		payload := schema.CreateTypeResponse{}
		payload.Type = schema.ConvertTypeToCreateResponse(data.Type)

		resp = response.Created(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}

func (h *Handler) UpdateType() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		typeBody := &schema.CreateTypeRequest{}
		err := json.NewDecoder(req.Body).Decode(typeBody)

		var resp *response.Response

		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		useCaseInput := usecase.UpdateTypeInput{}
		id := chi.URLParam(req, "id")

		idInt, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			resp := response.BadRequest(err, "Invalid ID")
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		useCaseInput.Type.ID = idInt
		useCaseInput.Type.Label = typeBody.Label

		_, err = h.useCase.GetTypeByLabel(req.Context(), usecase.GetTypeByLabelInput{
			Label: useCaseInput.Type.Label,
		})

		if err == nil {
			resp := response.BadRequest(erring.ErrLabelExists, erring.ErrLabelExists.Message)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		data, err := h.useCase.UpdateType(req.Context(), useCaseInput)
		if err != nil {
			resp := response.BadRequest(erring.ErrTypeNotFound, erring.ErrTypeNotFound.Message)
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		payload := schema.CreateTypeResponse{}
		payload.Type = schema.ConvertTypeToCreateResponse(data.Type)

		resp = response.Created(payload)
		rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck
	}
}

func (h *Handler) DeleteType() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		id := chi.URLParam(req, "id")

		idInt, err := strconv.ParseInt(id, 10, 32)
		if err != nil {
			resp := response.BadRequest(err, "Invalid ID")
			rest.SendJSON(rw, resp.Status, resp.Payload, resp.Headers) //nolint:errcheck

			return
		}

		err = h.useCase.DeleteType(req.Context(), usecase.DeleteTypeInput{
			ID: int32(idInt),
		})
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				resp := response.NotFound(erring.ErrTypeNotFound, erring.ErrTypeNotFound.Message)
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
