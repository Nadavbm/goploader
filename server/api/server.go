package api

import (
	"log"
	"net/http"
)

// Server web server
type Server struct {
	Mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		Mux: http.NewServeMux(),
	}
}

func (s *Server) StartServer() {
	s.Mux.HandleFunc("/upload", uploadHandler)

	log.Println("starting server in port 8080")
	if err := http.ListenAndServe(":8080", s.Mux); err != nil {
		log.Fatal(err)
	}
}
