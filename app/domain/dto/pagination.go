package dto

type ValidationPaginationFunc func(p Pagination) error

type Pagination struct {
	Page         int
	ItemsPerPage int
	SortBy       string
	SortType     string
}

func (p Pagination) Validate(validateFunc ValidationPaginationFunc) error {
	return validateFunc(p)
}
