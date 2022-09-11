package handlers

import (
	"net/http"

	"github.com/JPaulo-Moura/uuid-validator/api/uuid"
)

type ValidatorUUID interface {
	Validate(rw http.ResponseWriter, r *http.Request)
}

type Validator struct {
	Service uuid.UUID
}

func NewValidator(service uuid.UUID) Validator {
	return Validator{
		Service: service,
	}
}
