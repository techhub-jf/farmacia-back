package erring

var (
	ErrClientAlreadyExists  = NewAppError("client:could-not-create-client", "client already exists")
	ErrClientCpfInvalid     = NewAppError("client:invalid-cpf", "the CPF provided is invalid")
	ErrClientNullFields     = NewAppError("client:entity-can-not-contain-null-fields", "some fields provided are null")
)
