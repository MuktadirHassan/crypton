package api

import (
	"log"
	"net/http"
)

type Server struct {
	addr   string
	routes http.Handler
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:    s.addr,
		Handler: s.routes,
	}

	log.Println("Server is running on", s.addr)
	return server.ListenAndServe()
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, routes: Routes()}
}
