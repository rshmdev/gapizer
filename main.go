package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/rshmdev/gapizer/internal/generator"
	"github.com/rshmdev/gapizer/internal/parser"
)

func main() {
	configPath := flag.String("config", "configs/example.yml", "Caminho para o arquivo de configuração")
	outputDir := flag.String("output", "./output", "Diretório de saída do código gerado")
	flag.Parse()

	absConfigPath, err := filepath.Abs(*configPath)
	if err != nil {
		log.Fatalf("Erro ao converter caminho do arquivo de configuração: %v", err)
	}

	absOutputDir, err := filepath.Abs(*outputDir)
	if err != nil {
		log.Fatalf("Erro ao converter caminho do diretório de saída: %v", err)
	}

	fmt.Println("Caminho absoluto do arquivo de configuração:", absConfigPath)
	fmt.Println("Caminho absoluto do diretório de saída:", absOutputDir)

	if _, err := os.Stat(absConfigPath); os.IsNotExist(err) {
		log.Fatalf("Arquivo de configuração não encontrado: %s", absConfigPath)
	}

	config, err := parser.ParseConfig(absConfigPath)
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	err = generator.GenerateAPI(config, absOutputDir)
	if err != nil {
		log.Fatalf("Erro ao gerar API: %v", err)
	}

	fmt.Printf("API gerada com sucesso em: %s\n", absOutputDir)
}
