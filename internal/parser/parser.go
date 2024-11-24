package parser

import (
	"errors"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppName   string     `yaml:"app_name"`
	Port      int        `yaml:"port"`
	Database  string     `yaml:"database"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Name        string            `yaml:"name"`
	Method      string            `yaml:"method"`
	Request     map[string]string `yaml:"request,omitempty"`
	Response    map[string]string `yaml:"response,omitempty"`
	HandlerName string
}

func ParseConfig(filePath string) (*Config, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config

	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	if config.AppName == "" {
		return nil, errors.New("o campo 'app_name' é obrigatório")
	}
	if config.Port == 0 {
		return nil, errors.New("o campo 'port' é obrigatório")
	}
	if len(config.Endpoints) == 0 {
		return nil, errors.New("a lista de 'endpoints' não pode estar vazia")
	}

	for i, endpoint := range config.Endpoints {
		if endpoint.Name == "" || endpoint.Method == "" {
			return nil, errors.New("cada endpoint deve ter os campos 'name' e 'method'")
		}
		config.Endpoints[i].HandlerName = generateHandlerName(endpoint.Name, endpoint.Method)
	}

	return &config, nil
}

func generateHandlerName(name, method string) string {
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "{", "")
	name = strings.ReplaceAll(name, "}", "")
	return strings.Title(strings.ToLower(method)) + strings.Title(name)
}
