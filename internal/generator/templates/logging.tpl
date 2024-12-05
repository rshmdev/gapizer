package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

var logToFile *log.Logger

func InitLogger(output string, filePath string) {
	if output == "file" {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Erro ao abrir arquivo de log: %v", err)
		}
		logToFile = log.New(file, "", log.LstdFlags)
	} else {
		logToFile = log.Default()
	}
}

// LoggingMiddleware registra informações sobre as requisições
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Registrar a requisição
		log.Printf("Método: %s, Endpoint: %s, IP: %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Executar o próximo middleware/handler
		next.ServeHTTP(w, r)

		// Registrar o tempo de execução
		duration := time.Since(start)
		log.Printf("Endpoint: %s, Tempo de execução: %v", r.URL.Path, duration)
	})
}
