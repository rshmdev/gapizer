package handlers

import (
	"encoding/json"
	"net/http"
)


// Post_login é o handler para POST /login
func Post_login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	

	

	
	var request struct {
		
		Password string `json:"password"` 
		
		Username string `json:"username"` 
		
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Erro ao decodificar a requisição: "+err.Error(), http.StatusBadRequest)
		return
	}
	

	// Exemplo de resposta
	response := map[string]string{
		"message": "Handler para /login",
	}
	json.NewEncoder(w).Encode(response)
}

