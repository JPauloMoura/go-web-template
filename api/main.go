package main

import (
	"github.com/JPaulo-Moura/uuid-validator/api/config"
	"github.com/JPaulo-Moura/uuid-validator/api/handlers"
	"github.com/JPaulo-Moura/uuid-validator/api/server"
	"github.com/JPaulo-Moura/uuid-validator/api/uuid"
)

// @title API para validação de UUIDs
// @version 1.0.0
// @contact.name João Paulo Moura
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3002
// @BasePath /
func main() {
	handler := handlers.NewValidator(uuid.UUID{})
	s := server.NewServer(handler, config.GetPort())
	s.Start()
}
