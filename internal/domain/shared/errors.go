package shared

import "errors"

var (
	ErrInternal = errors.New("erro interno do sistema")

	ErrNotFound      = errors.New("recurso não encontrado")
	ErrAuthorization = errors.New("não autorizado")
	ErrValidation    = errors.New("erro de validação")
	ErrConflictError = errors.New("conflito de dados")
)
