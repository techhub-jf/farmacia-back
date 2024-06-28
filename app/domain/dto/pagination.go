package dto

type Pagination struct {
	Page         int
	ItemsPerPage int
	SortBy       string
	SortType     string
}

const (
	DefaultItemsPerPage = 10
)
