package schema

type Meta struct {
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
	TotalItems   int `json:"total_items"`
}

type PaginatedResponse[T any] struct {
	Items    []T  `json:"items"`
	Metadata Meta `json:"metadata"`
}

const (
	DefaultItemsPerPage = 10
)
