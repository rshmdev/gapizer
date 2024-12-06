package database

import (
	"database/sql"
	"log"
	{{ if eq .Type "sqlite" }}
	_ "github.com/mattn/go-sqlite3"
	{{ else if eq .Type "mysql" }}
	_ "github.com/go-sql-driver/mysql"
	{{ else if eq .Type "postgresql" }}
	_ "github.com/lib/pq"
	{{ end }}
)

var DB *sql.DB

func InitDatabase() {
	var err error

	{{ if eq .Type "sqlite" }}
	DB, err = sql.Open("sqlite3", "{{ .Name }}")
	{{ else if eq .Type "mysql" }}
	dsn := "{{ .Username }}:{{ .Password }}@tcp({{ .Host }}:{{ .Port }})/{{ .Name }}"
	DB, err = sql.Open("mysql", dsn)
	{{ else if eq .Type "postgresql" }}
	dsn := "host={{ .Host }} port={{ .Port }} user={{ .Username }} password={{ .Password }} dbname={{ .Name }} sslmode=disable"
	DB, err = sql.Open("postgres", dsn)
	{{ end }}

	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Erro ao verificar a conexão com o banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados {{ .Type }} estabelecida com sucesso.")
}
