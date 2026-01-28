package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
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
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "index.html")
}

func (h *Handler) HandleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.logger.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		h.logger.Printf("Error reading file: %v", err)
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	input, err := io.ReadAll(file)
	if err != nil {
		h.logger.Printf("Error reading file content: %v", err)
		http.Error(w, "Error reading file content", http.StatusInternalServerError)
		return
	}

	result, err := h.service.ConvertString(string(input))
	if err != nil {
		h.logger.Printf("Error converting string: %v", err)
		http.Error(w, "Error converting string", http.StatusInternalServerError)
		return
	}

	filename := "result_" + time.Now().Format("20060102_150405") + ".txt"
	if err := os.WriteFile(filename, []byte(result), 0644); err != nil {
		h.logger.Printf("Error saving result file: %v", err)

	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
