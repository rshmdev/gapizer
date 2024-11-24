package routes

import (
	"net/http"
	"{{ .Namespace }}/handlers" // Caminho dinâmico
)

// Register{{ .Resource | title }}Routes registra as rotas relacionadas a {{ .Resource }}
func Register{{ .Resource | title }}Routes(mux *http.ServeMux) {
	{{ range .Endpoints }}
	mux.HandleFunc("{{ .Name }}", handlers.{{ .HandlerName }})
	{{ end }}
}
