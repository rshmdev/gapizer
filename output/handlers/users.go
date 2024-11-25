package handlers

import (
	"encoding/json"
	"net/http"
)


// Get_users é o handler para GET /users
func Get_users(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	

	

	

	// Exemplo de resposta
	response := map[string]string{
		"message": "Handler para /users",
	}
	json.NewEncoder(w).Encode(response)
}

