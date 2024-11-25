package routes

import (
	"net/http"
	"output/handlers"
	"output/middleware"
)

// RegisterUsersRoutes registra as rotas relacionadas a users
func RegisterUsersRoutes(mux *http.ServeMux) {
	
	
	mux.Handle("/users", middleware.AuthMiddleware(http.HandlerFunc(handlers.Get_users)))
	
	
}
