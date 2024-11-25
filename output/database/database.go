package database

import (
	"database/sql"
	"log"
	
	_ "github.com/mattn/go-sqlite3"
	
)

var DB *sql.DB

func InitDatabase() {
	var err error

	
	DB, err = sql.Open("sqlite3", "./data.db")
	

	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados sqlite estabelecida com sucesso.")
}
