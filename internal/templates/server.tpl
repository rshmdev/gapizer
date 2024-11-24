package main

import (
	"log"
	"net/http"
	"{{ .Namespace }}/routes"
)

func main() {
	mux := http.NewServeMux()

	// Registrar rotas
	{{ range .Resources }}
	routes.Register{{ . | title }}Routes(mux)
	{{ end }}


	log.Println("Servidor rodando na porta {{ .Port }}...")
	log.Fatal(http.ListenAndServe(":{{ .Port }}", mux))
}
