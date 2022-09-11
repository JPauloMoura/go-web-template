package server

import (
	"log"
	"net/http"

	"github.com/JPaulo-Moura/uuid-validator/api/handlers"

	"github.com/rs/cors"
)

type IServer interface {
	Start()
}

type Server struct {
	HandlerUse handlers.ValidatorUUID
	Port       string
}

func NewServer(handler handlers.Validator, port string) IServer {
	return Server{
		HandlerUse: handler,
		Port:       port,
	}
}

func (s Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/validator", s.HandlerUse.Validate)
	routers := cors.AllowAll().Handler(mux)

	log.Println("Server running on port: ", s.Port)
	panic(http.ListenAndServe(":"+s.Port, routers))
}
