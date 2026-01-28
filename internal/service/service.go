package service

import (
	"os"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

type Service struct {
	converter morse.Converter
}

func NewService(converter morse.Converter) *Service {
	return &Service{
		converter: converter,
	}
}

func (s *Service) ConvertString(input string) (string, error) {
	// Преобразуем все символы в верхний регистр
	upperInput := strings.ToUpper(input)

	var mrs bool

	// Проверка, является ли строка кодом Морзе
	for _, char := range upperInput {
		if char != '.' && char != '-' && char != ' ' {
			mrs = false
			break
		}
		mrs = true
	}

	if mrs {
		// Если входная строка — это код Морзе, конвертируем в текст
		text := s.converter.ToText(upperInput)

		// Записываем в файл
		fileName := "morse_to_text_" + time.Now().Format("20060102_150405") + ".txt"
		os.WriteFile(fileName, []byte(text), 0644)

		return text, nil
	} else {
		// Если входная строка — это обычный текст, конвертируем в Морзе
		morseCode := s.converter.ToMorse(upperInput)

		// Записываем в файл
		fileName := "text_to_morse_" + time.Now().Format("20060102_150405") + ".txt"
		os.WriteFile(fileName, []byte(morseCode), 0644)

		return morseCode, nil
	}
}
