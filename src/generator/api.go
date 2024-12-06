package generator

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/rshmdev/gapizer/src/parser"
)

//go:embed templates/*
var templateFS embed.FS

func SanitizeFileName(name string) string {
	replacer := strings.NewReplacer("{", "_", "}", "_", "/", "_")
	return replacer.Replace(name)
}

func GenerateAPI(config *parser.Config, outputDir string) error {

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	if err := generateGoMod(outputDir); err != nil {
		return err
	}

	namespace := filepath.Base(outputDir)

	groupedRoutes := groupRoutesByResource(config.Endpoints)

	var resources []string
	for resource := range groupedRoutes {
		resources = append(resources, resource)
	}

	serverFile := filepath.Join(outputDir, "main.go")
	serverData := map[string]interface{}{
		"Namespace": namespace,
		"Port":      config.Port,
		"Resources": resources,
		"Logging":   config.Logging,
	}
	if err := generateFromTemplate("templates/server.tpl", serverFile, serverData); err != nil {
		return err
	}

	for resource, endpoints := range groupedRoutes {
		routeFile := filepath.Join(outputDir, "routes", fmt.Sprintf("%s.go", resource))
		handlerFile := filepath.Join(outputDir, "handlers", fmt.Sprintf("%s.go", resource))

		hasProtectedEndpoints := false
		for _, endpoint := range endpoints {
			if endpoint.Protected {
				hasProtectedEndpoints = true
				break
			}
		}

		data := map[string]interface{}{
			"Namespace":             namespace,
			"Resource":              resource,
			"Endpoints":             endpoints,
			"HasProtectedEndpoints": hasProtectedEndpoints,
		}

		if err := generateFromTemplate("templates/route.tpl", routeFile, data); err != nil {
			return err
		}
		if err := generateFromTemplate("templates/handler.tpl", handlerFile, data); err != nil {
			return err
		}
	}

	docsFile := filepath.Join(outputDir, "docs", "swagger.yaml")
	if err := generateFromTemplate("templates/swagger.tpl", docsFile, config); err != nil {
		return err
	}

	if config.Authentication != nil {
		middlewareFile := filepath.Join(outputDir, "middleware", "auth.go")
		if err := generateFromTemplate("templates/middleware.tpl", middlewareFile, config.Authentication); err != nil {
			return fmt.Errorf("erro ao gerar o middleware de autenticação: %w", err)
		}
	}

	if config.Logging != nil && config.Logging.Enabled {
		middlewareDir := filepath.Join(outputDir, "middleware")
		if err := os.MkdirAll(middlewareDir, 0755); err != nil {
			return fmt.Errorf("erro ao criar diretório de middleware: %w", err)
		}

		middlewareFile := filepath.Join(middlewareDir, "logging.go")
		if err := generateFromTemplate("templates/logging.tpl", middlewareFile, config.Logging); err != nil {
			return fmt.Errorf("erro ao gerar middleware de logging: %w", err)
		}
	}

	dbFile := filepath.Join(outputDir, "database", "database.go")
	if err := generateFromTemplate("templates/database.tpl", dbFile, config.Database); err != nil {
		return fmt.Errorf("erro ao gerar o arquivo de banco de dados: %w", err)
	}

	if err := updateGoMod(outputDir, config.Database.Type); err != nil {
		return fmt.Errorf("erro ao atualizar go.mod: %w", err)
	}

	if config.Database.Type == "sqlite" {
		fmt.Println(`
IMPORTANTE: Para usar SQLite, habilite o CGO_ENABLED ao rodar ou compilar o projeto:
  - Rodar: CGO_ENABLED=1 go run main.go
  - Compilar: CGO_ENABLED=1 go build -o myapp main.go
		`)
	}

	return nil
}

func updateGoMod(outputDir, dbType string) error {
	goModPath := filepath.Join(outputDir, "go.mod")

	// Dependências básicas
	dependencies := []struct {
		Module  string
		Version string
	}{
		// Dependência JWT
		{"github.com/golang-jwt/jwt/v4", "v4.5.0"},
	}

	// Dependências específicas para o banco de dados
	switch dbType {
	case "sqlite":
		dependencies = append(dependencies, struct {
			Module  string
			Version string
		}{"github.com/mattn/go-sqlite3", "v1.14.17"})
	case "mysql":
		dependencies = append(dependencies, struct {
			Module  string
			Version string
		}{"github.com/go-sql-driver/mysql", "v1.6.0"})
	case "postgresql":
		dependencies = append(dependencies, struct {
			Module  string
			Version string
		}{"github.com/lib/pq", "v1.10.5"})
	}

	// Ler o arquivo go.mod existente
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return fmt.Errorf("erro ao ler go.mod: %w", err)
	}

	// Adicionar dependências ao go.mod
	updatedContent := string(content)
	for _, dep := range dependencies {
		updatedContent += fmt.Sprintf("\nrequire %s %s\n", dep.Module, dep.Version)
	}

	if err := os.WriteFile(goModPath, []byte(updatedContent), 0644); err != nil {
		return fmt.Errorf("erro ao atualizar go.mod: %w", err)
	}

	// Rodar `go get` para instalar as dependências
	for _, dep := range dependencies {
		cmd := exec.Command("go", "get", dep.Module+"@"+dep.Version)
		cmd.Dir = outputDir
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("erro ao executar 'go get' para %s: %w", dep.Module, err)
		}
	}

	return nil
}

func generateFromTemplate(templatePath, outputPath string, data interface{}) error {
	funcMap := template.FuncMap{
		"lower": strings.ToLower,
		"title": strings.Title,
	}

	// Carregar o conteúdo do template a partir do embed
	templateContent, err := templateFS.ReadFile(templatePath)
	if err != nil {
		return fmt.Errorf("erro ao carregar template embutido: %w", err)
	}

	// Parse o conteúdo do template
	tmpl, err := template.New(filepath.Base(templatePath)).Funcs(funcMap).Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("erro ao parsear template: %w", err)
	}

	// Criar o diretório do arquivo de saída
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório do arquivo de saída: %w", err)
	}

	// Criar o arquivo de saída
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de saída: %w", err)
	}
	defer file.Close()

	// Executar o template e escrever no arquivo de saída
	return tmpl.Execute(file, data)
}

func groupRoutesByResource(endpoints []parser.Endpoint) map[string][]parser.Endpoint {
	grouped := make(map[string][]parser.Endpoint)
	for _, endpoint := range endpoints {
		resource := getResourceName(endpoint.Name)
		grouped[resource] = append(grouped[resource], endpoint)
	}
	return grouped
}

func generateGoMod(outputDir string) error {
	moduleName := filepath.Base(outputDir)
	goModContent := fmt.Sprintf(`module %s

go 1.20
`, moduleName)

	goModPath := filepath.Join(outputDir, "go.mod")
	if err := os.WriteFile(goModPath, []byte(goModContent), 0644); err != nil {
		return fmt.Errorf("erro ao criar go.mod: %w", err)
	}

	return nil
}

func getResourceName(endpoint string) string {
	if endpoint[0] == '/' {
		endpoint = endpoint[1:]
	}
	parts := strings.Split(endpoint, "/")
	return parts[0]
}
