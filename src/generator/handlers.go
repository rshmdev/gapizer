package generator

import (
	"fmt"
	"os"
	"text/template"
)

type Handler struct {
	Name string
}

func GenerateHandlers(outputDir string, handlers []Handler) error {
	handlerTemplate := `
package handlers

import (
	"net/http"
)

{{ range . }}
func {{ .Name }}Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Handler for {{ .Name }}"))
}
{{ end }}
`

	tmpl, err := template.New("handlers").Parse(handlerTemplate)
	if err != nil {
		return err
	}

	handlerFile := fmt.Sprintf("%s/handlers/handlers.go", outputDir)
	file, err := os.Create(handlerFile)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, handlers)
}
