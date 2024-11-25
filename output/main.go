package main

import (
	"log"
	"net/http"
	"output/routes"
	"output/database"
    "output/middleware"

)

func main() {

	
	middleware.InitLogger("console", "logs/server.log")
	

    database.InitDatabase()

	mux := http.NewServeMux()

   	
	// Aplicar middleware de logging globalmente
	loggingMux := middleware.LoggingMiddleware(mux)
	


	// Registrar rotas
	
	routes.RegisterUsersRoutes(mux)
	
	routes.RegisterLoginRoutes(mux)
	


	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", loggingMux))
}
