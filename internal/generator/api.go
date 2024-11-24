package generator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gapizer/internal/parser"
)

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
	}
	if err := generateFromTemplate("server.tpl", serverFile, serverData); err != nil {
		return err
	}

	for resource, endpoints := range groupedRoutes {
		routeFile := filepath.Join(outputDir, "routes", fmt.Sprintf("%s.go", resource))
		handlerFile := filepath.Join(outputDir, "handlers", fmt.Sprintf("%s.go", resource))

		data := map[string]interface{}{
			"Namespace": namespace,
			"Resource":  resource,
			"Endpoints": endpoints,
		}

		if err := generateFromTemplate("route.tpl", routeFile, data); err != nil {
			return err
		}
		if err := generateFromTemplate("handler.tpl", handlerFile, data); err != nil {
			return err
		}
	}

	docsFile := filepath.Join(outputDir, "docs", "swagger.yaml")
	if err := generateFromTemplate("swagger.tpl", docsFile, config); err != nil {
		return err
	}

	return nil
}

func generateFromTemplate(templatePath, outputPath string, data interface{}) error {
	funcMap := template.FuncMap{
		"lower": strings.ToLower,
		"title": strings.Title,
	}

	tmplPath := filepath.Join("internal", "templates", templatePath)

	tmpl, err := template.New(filepath.Base(templatePath)).Funcs(funcMap).ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("erro ao carregar template: %w", err)
	}

	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório do arquivo de saída: %w", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de saída: %w", err)
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}

func copySwaggerUI(outputDir string) error {
	srcDir := "internal/templates/swagger-ui"
	destDir := filepath.Join(outputDir, "swagger-ui")

	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(srcDir, path)
		destPath := filepath.Join(destDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}

		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(destPath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		return err
	})
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
	goModContent := fmt.Sprintf("module %s\n\ngo 1.20\n", moduleName)

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
