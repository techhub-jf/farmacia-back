package erring

var (
	ErrLabelExists  = NewAppError("type:exists", "type already exists.")
	ErrTypeNotFound = NewAppError("type:notFound", "type not found.")
)
