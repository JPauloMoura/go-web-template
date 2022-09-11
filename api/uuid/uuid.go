package uuid

import (
	"log"

	"github.com/google/uuid"
)

type UUID struct {
	UID     string       `json:"uid"`
	Version uuid.Version `json:"version"`
	IsValid bool         `json:"valid"`
}

func (u UUID) Parse(uid string) (*UUID, error) {
	uuid, err := uuid.Parse(uid)
	if err != nil {
		log.Printf("method:%s, error:%s", "UUID.Parse", errInvalidUUID)
		return nil, errInvalidUUID
	}

	return &UUID{
		UID:     uid,
		Version: uuid.Version(),
		IsValid: true,
	}, nil
}
