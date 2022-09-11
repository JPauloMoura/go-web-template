package main

import (
	"github.com/JPaulo-Moura/uuid-validator/api/config"
	"github.com/JPaulo-Moura/uuid-validator/api/handlers"
	"github.com/JPaulo-Moura/uuid-validator/api/server"
	"github.com/JPaulo-Moura/uuid-validator/api/uuid"
)

func main() {
	handler := handlers.NewValidator(uuid.UUID{})
	s := server.NewServer(handler, config.GetPort())
	s.Start()
}
