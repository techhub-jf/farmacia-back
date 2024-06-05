package erring

var (
	ErrLoginUnauthorized    = NewAppError("login:unauthorized", "Email ou senha incorretos!")
	ErrLoginUserNotFound    = NewAppError("login:user-not-found", "Usuário não encontrado!")
	ErrLoginTokenNotCreated = NewAppError("login:token-not-created", "Erro interno.")
)
