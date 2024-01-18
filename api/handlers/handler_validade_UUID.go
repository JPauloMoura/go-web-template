package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JPaulo-Moura/uuid-validator/api/uuid"
)

// swagger:response validationError
type ValidatorResponse struct {
	Data  *uuid.UUID `json:"data"`
	Error string     `json:"error"`
}

// @Summary Verifica se a valor informado é um uuid valido
// @Description Verifica se a valor informado é um uuid valido
// @Tags Validate
// @Accept  json
// @Produce  json
// @Param uid query string true "6457d5dc-6a4b-409f-972e-f8bb8f9f9f67"
// @Success 200 {object} ValidatorResponse
// @Router /validator [get]
func (h Validator) Validate(rw http.ResponseWriter, r *http.Request) {
	var resp ValidatorResponse
	uid := r.URL.Query().Get("uid")
	if uid == "" {
		resp.Error = errUIDIsEmpty.Error()
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(resp)
		log.Printf("method:%s, error:%s", "Validator.Validate", resp.Error)
		return
	}

	uuid, err := h.Service.Parse(uid)
	if err != nil {
		resp.Error = err.Error()
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(resp)
		log.Printf("method:%s, error:%s", "Validator.Validate", resp.Error)
		return
	}

	resp.Data = uuid
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(resp)
	log.Printf("method:%s, sucess:%v", "Validator.Validate", *resp.Data)
}
