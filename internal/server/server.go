package server

import (
	"log"
	"net/http"
	"time"
)

type handler interface {
	HandleRoot(w http.ResponseWriter, r *http.Request)
	HandleUpload(w http.ResponseWriter, r *http.Request)
}

type Server struct {
	logger *log.Logger
	server *http.Server
}

func NewRouter(
	logger *log.Logger,
	addr string,
	port string,
	writeTimeout int,
	readTimeout int,
	idleTimeout int,
	handler handler,
) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HandleRoot)
	mux.HandleFunc("/upload", handler.HandleUpload)
	server := &http.Server{
		Addr:         addr + ":" + port,
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		IdleTimeout:  time.Duration(idleTimeout) * time.Second,
	}
	return &Server{
		server: server,
		logger: logger,
	}
}

func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s\n", s.server.Addr)
	return s.server.ListenAndServe()
}
