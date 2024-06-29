package handler

import (
	"net/url"
	"strconv"

	"github.com/techhub-jf/farmacia-back/app/gateway/api/handler/schema"
)

// The readString() helper returns a string value from the query string, or the provided
// default value if no matching key could be found.
func (h *Handler) readString(qs url.Values, key string, defaultValue string) string {
	str := qs.Get(key)

	if str == "" {
		return defaultValue
	}

	return str
}

// The readInt() helper reads a string value from the query string and converts it to an
// integer before returning. If no matching key could be found it returns the provided
// default value. If the value couldn't be converted to an integer, then we return
// the default value.
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

func (h *Handler) getPaginationParams(qs url.Values, input *schema.Pagination) {
	input.Page = h.readInt(qs, "page", 1)
	input.SortBy = h.readString(qs, "sortBy", "id")
	input.SortType = h.readString(qs, "sortType", "ASC")
	input.ItemsPerPage = h.readInt(qs, "itemsPerPage", schema.DefaultItemsPerPage)
}
