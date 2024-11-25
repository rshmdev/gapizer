package routes

import (
	"net/http"
	"output/handlers"
	
)

// RegisterLoginRoutes registra as rotas relacionadas a login
func RegisterLoginRoutes(mux *http.ServeMux) {
	
	
	mux.HandleFunc("/login", handlers.Post_login)
	
	
}
