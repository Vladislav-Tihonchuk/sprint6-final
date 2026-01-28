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
	// Создаем логгер
	logger := log.New(os.Stdout, "", 0)

	// Создаем конвертер Морзе
	converter := morse.NewConverter(morse.DefaultMorse)

	// Создаем сервис с конвертером
	mainService := service.NewService(converter)

	// Создаем обработчик
	handler := handlers.New(logger, mainService)

	// Создаем и запускаем сервер
	serv := server.NewRouter(logger, "localhost", "8080", 5, 10, 15, handler)
	serv.Start()
}
