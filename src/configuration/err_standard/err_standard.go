package errstandard

import "net/http"

// estrutura de erro
type ErrorStandard struct {
	//messagem de erro
	Message string `json:"message"`
	// mensagem do tipo de erro: eEX: BadRequest 400 e tals
	Err string `json:"errror"`
	//codifo em si do erro
	Code int `json:"code_error"`
	//causas de erro  podemos também configurar caso o causes venha vazio para nçao
	//aparecer nojson com a propriedade : `json:causes,omitempty`
	Causes []Cause `json:"causes"`
}

// estrutura de cxusa de erro
type Cause struct {
	Field   string `json:"field"`
	Message string `json:"warning_message"`
}

// cria uma intercace para satisfacer algumas blibliotecas que precisam d eum erro por conta de seguir uma intercvasde
func (err ErrorStandard) Error() string {
	return err.Message
}

// fnção que constroi/fabrica um novo errorStandard
// PARAMETROS P/ CRIAR UM ERRORSTANDARD:                                   RETORA A LOCALIXAÇÃO DE MEMORIA DO ERROR
func NewErrorStandard(Message string, Err string, Code int, Causes []Cause) *ErrorStandard {
	// retorna o endereço de memoria que vai apontar agora para o erro que a função
	//cria
	return &ErrorStandard{
		Message: Message,
		Err:     Err,
		Code:    Code,
		Causes:  Causes,
	}
}

// templete de erro badrequests
func NewErrorBadRequest(Message string) *ErrorStandard {
	return &ErrorStandard{
		Message: Message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

// templete de badrequestc com validação
func NewErrorBadRequestValidation(Message string, Causes []Cause) *ErrorStandard {
	return &ErrorStandard{
		Message: Message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
		Causes:  Causes,
	}
}

//templete de server nternal error

func NewErrorInternalServer(Message string) *ErrorStandard {
	return &ErrorStandard{
		Message: Message,
		Err:     "Server Internal Error",
		Code:    http.StatusInternalServerError,
	}
}

// templete erroUnauthorized
func NewErrorUnauthorized(Message string) *ErrorStandard {
	return &ErrorStandard{
		Message: Message,
		Err:     "Unauthorized",
		Code:    http.StatusUnauthorized,
	}
}
