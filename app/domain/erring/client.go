package erring

var (
	ErrClientAlreadyExists  = NewAppError("client:could-not-create-client", "client already exists")
	ErrClientCpfInvalid     = NewAppError("client:invalid-cpf", "the CPF provided is invalid")
	ErrClientEmptyFields     = NewAppError("client:entity-can-not-contain-empty-fields", "some or all fields provided are empty")
)
