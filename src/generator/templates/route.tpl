package routes

import (
	"net/http"
	"{{ .Namespace }}/handlers"
	{{ if .HasProtectedEndpoints }}"{{ .Namespace }}/middleware"{{ end }}
)

// Register{{ .Resource | title }}Routes registra as rotas relacionadas a {{ .Resource }}
func Register{{ .Resource | title }}Routes(mux *http.ServeMux) {
	{{ range .Endpoints }}
	{{ if .Protected }}
	mux.Handle("{{ .Name }}", middleware.AuthMiddleware(http.HandlerFunc(handlers.{{ .HandlerName }})))
	{{ else }}
	mux.HandleFunc("{{ .Name }}", handlers.{{ .HandlerName }})
	{{ end }}
	{{ end }}
}
