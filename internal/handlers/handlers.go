package handlers

import (
	"log"
	"net/http"
	"path/filepath"
)

type service interface {
	ConvertString(input string) (string, error)
}

type Handler struct {
	logger  *log.Logger
	service service
}

func New(logger *log.Logger, service service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {
	filepath := filepath.Join("../index.html")
	http.ServeFile(w, r, filepath)
}

func (h *Handler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10000000000); err != nil {
		h.logger.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	input := r.MultipartForm.File["myFile"][0]
	result, err := h.service.ConvertString(input)
	if err != nil {
		h.logger.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "test/html")
	w.Write([]byte(result))
}
