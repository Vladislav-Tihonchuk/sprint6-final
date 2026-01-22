package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Простой логгер
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Создаем сервер
	serv := server.New(logger)

	// Регистрируем маршруты
	serv.RegisterRoutes()

	// Запускаем
	logger.Println("Сервер запущен на http://localhost:8080")
	if err := serv.Start(); err != nil {
		logger.Fatal("Ошибка сервера:", err)
	}
}
