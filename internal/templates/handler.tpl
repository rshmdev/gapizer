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

	// Exemplo de resposta
	response := map[string]string{
		"message": "Handler para {{ .Name }}",
	}
	json.NewEncoder(w).Encode(response)
}
{{ end }}
