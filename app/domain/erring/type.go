package erring

var (
	ErrLabelExists = NewAppError("type:exists", "type already exists.")
)
