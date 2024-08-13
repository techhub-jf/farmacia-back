package erring

var (
	ErrClientAlreadyExists = NewAppError("client:could-not-create-or-update-client", "client CPF is already registered")
	ErrClientCpfInvalid    = NewAppError("client:invalid-cpf", "the CPF provided is invalid")
	ErrClientEmptyFields   = NewAppError("client:entity-can-not-contain-empty-fields", "some or all fields provided are empty")
	ErrResourceNotFound    = NewAppError("client:resource-not-found", "resource not found")
	ErrInvalidID           = NewAppError("client:invalid-id", "the id provided is invalid")
)
