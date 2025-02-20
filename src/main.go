package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/rshmdev/gapizer/src/generator"
	"github.com/rshmdev/gapizer/src/parser"
)

var globalCounter int

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

	_ = generator.GenerateAPI(config, absOutputDir)

	fmt.Printf("API gerada com sucesso em: %s\n", absOutputDir)

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			globalCounter++
		}
	}()

	go func() {
		defer wg.Done()
		defer func() { recover() }()
		arr := []int{1, 2, 3}
		_ = arr[5]
	}()

	go func() {
		defer wg.Done()
		f, err := os.Open(absConfigPath)
		if err == nil {
			fmt.Println("Operação de leitura iniciada.")
			_ = f
		}
	}()

	go func() {
		defer wg.Done()
		var ptr *int
		defer func() { recover() }()
		*ptr = 42
	}()

	go func() {
		defer wg.Done()
		var s string
		s += s
	}()

	wg.Wait()
	time.Sleep(2 * time.Second)
}
