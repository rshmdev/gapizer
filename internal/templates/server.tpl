package main

import (
	"log"
	"net/http"
	"{{ .Namespace }}/routes"
	"{{ .Namespace }}/database"
    {{ if .Logging.Enabled }}"{{ .Namespace }}/middleware"{{ end }}

)

func main() {

	{{ if .Logging.Enabled }}
	middleware.InitLogger("{{ .Logging.Output }}", "{{ .Logging.FilePath }}")
	{{ end }}

    database.InitDatabase()

	mux := http.NewServeMux()

   	{{ if .Logging.Enabled }}
	// Aplicar middleware de logging globalmente
	loggingMux := middleware.LoggingMiddleware(mux)
	{{ else }}
	loggingMux := mux
	{{ end }}


	// Registrar rotas
	{{ range .Resources }}
	routes.Register{{ . | title }}Routes(mux)
	{{ end }}


	log.Println("Servidor rodando na porta {{ .Port }}...")
	log.Fatal(http.ListenAndServe(":{{ .Port }}", loggingMux))
}
