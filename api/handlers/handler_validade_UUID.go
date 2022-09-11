package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JPaulo-Moura/uuid-validator/api/uuid"
)

type ValidatorResponse struct {
	Data  *uuid.UUID `json:"data"`
	Error string     `json:"error"`
}

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
