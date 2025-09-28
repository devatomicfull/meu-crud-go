package rest_erro

import "net/http"

// RestErro representa um erro padronizado que será retornado em respostas HTTP.
// Cada campo possui uma tag JSON para que os nomes no JSON sejam controlados.
// Passando o Objeto para JSON (Exportavel para Json usando Apostrofo)
type RestErro struct {
	// Message é a mensagem de erro amigável para o cliente.
	// A tag `json:"message"` garante que no JSON o campo apareça como "message".
	Message string `json:"message"`

	// Erro é um identificador técnico do erro (ex.: bad_request, not_found).
	Erro string `json:"erro"`

	// Code é o código de status HTTP associado ao erro.
	Code int `json:"code"`

	// Causes é uma lista de causas específicas do erro, útil para validações campo a campo.
	Causes []Causes `json:"causes"`
}

// Causes representa as causas específicas de um erro de validação.
// Útil para detalhar problemas em campos específicos.
type Causes struct {
	// Field é o nome do campo que gerou o erro.
	Field string `json:"field"`

	// Message é a mensagem de erro específica para esse campo.
	Message string `json:"message"`
}

// Error implementa a interface de erro padrão do Go (error).
// Sempre que esse objeto for usado como erro, essa string será retornada.
func (r *RestErro) Error() string {
	return r.Message
}

// NewRestErro é um construtor genérico para criar um RestErro personalizado.
// Note o uso de &RestErro{...} retornando um *RestErro:
// - O & cria um ponteiro para a struct (endereço na memória).
// - O * no tipo *RestErro indica que a função retorna um ponteiro para RestErro.
// Isso evita cópias desnecessárias e permite modificar o objeto original.
func NewRestErro(message string, erro string, code int, causes []Causes) *RestErro {
	return &RestErro{
		Message: message,
		Erro:    erro,
		Code:    code,
		Causes:  causes,
	}
}

// NewBadRequestError cria um erro com status 400 (Bad Request).
// Não recebe causas específicas.
func NewBadRequestError(message string) *RestErro {
	return &RestErro{
		Message: message,
		Erro:    "bad_request",
		Code:    http.StatusBadRequest, // Código HTTP 400
	}
}

// NewBadRequestValidationError cria um erro 400 (Bad Request) com detalhes
// sobre os campos inválidos (causas).
func NewBadRequestValidationError(message string, causes []Causes) *RestErro {
	return &RestErro{
		Message: message,
		Erro:    "bad_request",
		Code:    http.StatusBadRequest, // Código HTTP 400
		Causes:  causes,
	}
}

// NewInternalServerError cria um erro 500 (Internal Server Error).
func NewInternalServerError(message string) *RestErro {
	return &RestErro{
		Message: message,
		Erro:    "internal_server_error",
		Code:    http.StatusInternalServerError, // Código HTTP 500
	}
}

// NewNotFoundError cria um erro 404 (Not Found).
func NewNotFoundError(message string) *RestErro {
	return &RestErro{
		Message: message,
		Erro:    "not_found",
		Code:    http.StatusNotFound, // Código HTTP 404
	}
}

// NewConflictError cria um erro 409 (Conflict).
func NewConflictError(message string) *RestErro {
	return &RestErro{
		Message: message,
		Erro:    "conflict",
		Code:    http.StatusConflict, // Código HTTP 409
	}
}

// NewForbiddenError cria um erro 403 (Forbidden).
func NewForbiddenError(message string) *RestErro {
	return &RestErro{
		Message: message,
		Erro:    "forbidden",
		Code:    http.StatusForbidden, // Código HTTP 403
	}
}
