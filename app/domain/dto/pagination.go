package dto

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
