package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type Route struct {
	Name   string
	Method string
}

func GenerateRoutes(outputDir string, routes []Route) error {
	routeTemplate := `
package routes

import "net/http"

// RegisterRoutes registra todas as rotas da aplicação
func RegisterRoutes(mux *http.ServeMux) {
	{{ range . }}
	mux.HandleFunc("{{ .Name }}", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "{{ .Method }}" {
			w.Write([]byte("{{ .Method }} {{ .Name }} handler"))
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	{{ end }}
}
`

	routesDir := filepath.Join(outputDir, "routes")
	if err := os.MkdirAll(routesDir, 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório de rotas: %w", err)
	}

	for _, route := range routes {
		sanitizedName := SanitizeFileName(route.Name)
		routeFile := filepath.Join(routesDir, sanitizedName+".go")

		file, err := os.Create(routeFile)
		if err != nil {
			return fmt.Errorf("erro ao criar arquivo de rota %s: %w", routeFile, err)
		}
		defer file.Close()

		tmpl, err := template.New("route").Parse(routeTemplate)
		if err != nil {
			return fmt.Errorf("erro ao processar template de rota: %w", err)
		}

		if err := tmpl.Execute(file, route); err != nil {
			return fmt.Errorf("erro ao escrever template para %s: %w", routeFile, err)
		}
	}

	return nil
}
