package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	converter := morse.NewConverter(morse.DefaultMorse)
	mainService := service.NewService(converter)
	handl := handlers.New(logger, mainService)
	serv := server.NewRouter(logger, "localhost", "8080", 5, 10, 15, handl)
	serv.Start()
}
