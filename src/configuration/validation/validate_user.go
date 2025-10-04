package validation

import (
	"encoding/json"
	"errors"

	"github.com/devatomicfull/meu-crud-go/src/configuration/rest_erro"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"                                   // Pacote de tradução para inglês
	ut "github.com/go-playground/universal-translator"                      // Gerencia múltiplos idiomas
	"github.com/go-playground/validator/v10"                                // Validador principal
	enTranslations "github.com/go-playground/validator/v10/translations/en" // Traduções padrão do validator
)

// Variáveis globais
var (
	// ValidatorInstance é a instância principal do validador usada em todo o sistema.
	ValidatorInstance = validator.New()

	// Translator é o tradutor de mensagens de validação (por padrão, em inglês).
	Translator ut.Translator
)

// Constante para definir o idioma padrão do sistema
const (
	LocaleEN = "en"
)

// init é executado automaticamente ao iniciar o programa.
// Ele configura o validador padrão do Gin com suporte à tradução de mensagens.
func init() {
	// Verifica se o validador interno do Gin pode ser convertido para *validator.Validate
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Cria o tradutor base para o idioma inglês
		enLocale := en.New()
		universalTranslator := ut.New(enLocale, enLocale)

		// Obtém o tradutor para o idioma inglês
		translator, _ := universalTranslator.GetTranslator(LocaleEN)

		// Registra as traduções padrão (mensagens de erro amigáveis)
		enTranslations.RegisterDefaultTranslations(val, translator)

		// Armazena as referências globais
		ValidatorInstance = val
		Translator = translator
	}
}

// ValidateUserError
// Trata e converte erros de validação (do validator ou JSON) em um formato padronizado para resposta HTTP.
//
// Parâmetros:
//   - validationErr: erro retornado durante a validação ou decodificação de dados.
//
// Retorno:
//   - *rest_erro.RestErro: objeto de erro padronizado, pronto para ser retornado na API.
//
// A função identifica três tipos de erro principais:
//
//	Erro de tipo incorreto no JSON (ex: string em vez de número)
//	Erros de validação dos campos com mensagens traduzidas
//	Qualquer outro erro genérico de parsing/conversão
func ValidateUserError(validationErr error) *rest_erro.RestErro {
	// Verifica se o erro é de tipo incorreto no JSON (ex: "expected string but got number")
	var jsonTypeError *json.UnmarshalTypeError

	// Verifica se o erro é de validação de campos (ex: "email obrigatório")
	var validationErrors validator.ValidationErrors

	// Caso 1: Erro de tipo incorreto no JSON
	if errors.As(validationErr, &jsonTypeError) {
		return rest_erro.NewBadRequestError("Tipo de campo inválido. Verifique os dados enviados.")
	} else if errors.As(validationErr, &validationErrors) { // Caso 2: Erros de validação (campos que não atendem às regras do validator)
		errorsCauses := []rest_erro.Causes{}

		// Para cada erro de validação, cria uma causa com mensagem traduzida e o nome do campo
		for _, err := range validationErr.(validator.ValidationErrors) {
			cause := rest_erro.Causes{
				Field:   err.Field(),
				Message: err.Translate(Translator), // Usa o tradutor configurado no init()
			}
			errorsCauses = append(errorsCauses, cause)
		}

		// Retorna um erro padronizado com todas as causas
		return rest_erro.NewBadRequestValidationError(
			"Alguns campos são inválidos. Corrija os erros e tente novamente.",
			errorsCauses,
		)
	} else {
		// Caso 3: Qualquer outro tipo de erro inesperado
		return rest_erro.NewBadRequestError("Erro ao tentar converter ou validar os campos.")
	}
}
