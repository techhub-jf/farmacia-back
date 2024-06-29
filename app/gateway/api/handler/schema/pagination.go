package schema

type Meta struct {
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
	TotalItems   int `json:"total_items"`
}

type PaginationResponse[T any] struct {
	Items    []T  `json:"items"`
	Metadata Meta `json:"metadata"`
}

type ValidationFunc func(p Pagination) error

type Pagination struct {
	Page         int
	ItemsPerPage int
	SortBy       string
	SortType     string
}

func (p Pagination) Validate(validateFunc ValidationFunc) error {
	return validateFunc(p)
}

const (
	DefaultItemsPerPage = 10
)
