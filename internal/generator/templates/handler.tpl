package handlers

import (
	"encoding/json"
	"net/http"
)

{{ range .Endpoints }}
// {{ .HandlerName }} é o handler para {{ .Method }} {{ .Name }}
func {{ .HandlerName }}(w http.ResponseWriter, r *http.Request) {
	if r.Method != "{{ .Method }}" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	{{ if .Headers }}
	// Validar headers
	{{ range $key, $value := .Headers }}
	{{ $key | title }} := r.Header.Get("{{ $key }}")
	if {{ $key | title }} == "" {
		http.Error(w, "Header '{{ $key }}' é obrigatório", http.StatusBadRequest)
		return
	}
	{{ end }}
	{{ end }}

	{{ if .QueryParams }}
	// Validar query params
	query := r.URL.Query()
	{{ range $key, $value := .QueryParams }}
	{{ $key | title }} := query.Get("{{ $key }}")
	if {{ $key | title }} == "" {
		http.Error(w, "Query param '{{ $key }}' é obrigatório", http.StatusBadRequest)
		return
	}
	{{ end }}
	{{ end }}

	{{ if .Request }}
	var request struct {
		{{ range $key, $value := .Request }}
		{{ $key | title }} {{ $value }} `json:"{{ $key }}"` 
		{{ end }}
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Erro ao decodificar a requisição: "+err.Error(), http.StatusBadRequest)
		return
	}
	{{ end }}

	// Exemplo de resposta
	response := map[string]string{
		"message": "Handler para {{ .Name }}",
	}
	json.NewEncoder(w).Encode(response)
}
{{ end }}
