package handler

import (
	"net/url"
	"strconv"

	"github.com/techhub-jf/farmacia-back/app/domain/dto"
	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

func (h *Handler) readString(qs url.Values, key string, defaultValue string) string {
	str := qs.Get(key)

	if str == "" {
		return defaultValue
	}

	return str
}

func (h *Handler) readInt(qs url.Values, key string, defaultValue int) int {
	str := qs.Get(key)

	if str == "" {
		return defaultValue
	}

	convertedNumber, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}

	return convertedNumber
}

func (h *Handler) getPaginationParams(qs url.Values, input *dto.Pagination) {
	input.Page = h.readInt(qs, "page", 1)
	input.SortBy = h.readString(qs, "sortBy", "id")
	input.SortType = h.readString(qs, "sortType", "ASC")
	input.ItemsPerPage = h.readInt(qs, "itemsPerPage", schema.DefaultItemsPerPage)
}
