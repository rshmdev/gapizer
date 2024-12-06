openapi: 3.0.0
info:
  title: {{ .AppName }}
  version: 1.0.0
paths:
  {{ range .Endpoints }}
  {{ .Name | lower }}:
    {{ .Method | lower }}:
      summary: "Handler para {{ .Name }}"
      responses:
        "200":
          description: "Resposta bem-sucedida"
  {{ end }}
