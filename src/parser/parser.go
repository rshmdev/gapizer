package parser

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type LoggingConfig struct {
	Enabled  bool   `yaml:"enabled"`
	Output   string `yaml:"output"`
	FilePath string `yaml:"file_path,omitempty"`
}
type AuthenticationConfig struct {
	Type                   string `yaml:"type"`
	Secret                 string `yaml:"secret"`
	TokenExpirationMinutes int    `yaml:"token_expiration_minutes"`
}

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host,omitempty"`
	Port     int    `yaml:"port,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Name     string `yaml:"name"`
}

type Config struct {
	AppName        string                `yaml:"app_name"`
	Port           int                   `yaml:"port"`
	Database       DatabaseConfig        `yaml:"database"`
	Authentication *AuthenticationConfig `yaml:"authentication,omitempty"`
	Logging        *LoggingConfig        `yaml:"logging,omitempty"`
	Endpoints      []Endpoint            `yaml:"endpoints"`
}

type Endpoint struct {
	Name        string            `yaml:"name"`
	Method      string            `yaml:"method"`
	Request     map[string]string `yaml:"request,omitempty"`
	Response    map[string]string `yaml:"response,omitempty"`
	Headers     map[string]string `yaml:"headers,omitempty"`
	QueryParams map[string]string `yaml:"query_params,omitempty"`
	HandlerName string
	Protected   bool `yaml:"protected,omitempty"` // Novo campo para endpoints protegidos

}

func ValidateRequest(request map[string]string) error {
	validTypes := []string{"string", "int", "float64", "bool"}

	for key, typ := range request {
		isValidType := false
		for _, validType := range validTypes {
			if typ == validType {
				isValidType = true
				break
			}
		}
		if !isValidType {
			return fmt.Errorf("o campo '%s' possui um tipo inválido: '%s'. Tipos válidos são: %v", key, typ, validTypes)
		}
	}

	return nil
}

func ValidateLoggingConfig(logging *LoggingConfig) error {
	if !logging.Enabled {
		return nil
	}

	if logging.Output != "console" && logging.Output != "file" {
		return fmt.Errorf("o campo 'output' deve ser 'console' ou 'file'")
	}

	if logging.Output == "file" && logging.FilePath == "" {
		return errors.New("o campo 'file_path' é obrigatório quando 'output' é 'file'")
	}

	return nil
}

func ValidateAuthenticationConfig(auth *AuthenticationConfig) error {
	if auth.Type != "jwt" {
		return fmt.Errorf("tipo de autenticação '%s' não é suportado", auth.Type)
	}

	if auth.Secret == "" {
		return errors.New("o campo 'secret' é obrigatório para autenticação JWT")
	}

	if auth.TokenExpirationMinutes <= 0 {
		return errors.New("o campo 'token_expiration_minutes' deve ser maior que zero")
	}

	return nil
}

func ParseConfig(filePath string) (*Config, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo de configuração: %w", err)
	}

	var config Config

	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("erro ao fazer unmarshal do YAML: %w", err)
	}

	if config.AppName == "" {
		return nil, errors.New("o campo 'app_name' é obrigatório")
	}
	if config.Port == 0 {
		return nil, errors.New("o campo 'port' é obrigatório")
	}

	if config.Authentication != nil {
		if err := ValidateAuthenticationConfig(config.Authentication); err != nil {
			return nil, fmt.Errorf("erro na configuração de autenticação: %w", err)
		}
	}

	if config.Logging != nil {
		if err := ValidateLoggingConfig(config.Logging); err != nil {
			return nil, fmt.Errorf("erro na configuração de logs: %w", err)
		}
	}

	if err := ValidateDatabaseConfig(config.Database); err != nil {
		return nil, fmt.Errorf("erro na configuração do banco de dados: %w", err)
	}

	if len(config.Endpoints) == 0 {
		return nil, errors.New("a lista de 'endpoints' não pode estar vazia")
	}

	hasProtectedEndpoints := false
	for i, endpoint := range config.Endpoints {
		if endpoint.Name == "" || endpoint.Method == "" {
			return nil, fmt.Errorf("o endpoint na posição %d deve ter os campos 'name' e 'method'", i)
		}

		config.Endpoints[i].HandlerName = generateHandlerName(endpoint.Name, endpoint.Method)

		if endpoint.Protected {
			hasProtectedEndpoints = true
		}

		if len(endpoint.Request) > 0 {
			if err := ValidateRequest(endpoint.Request); err != nil {
				return nil, fmt.Errorf("erro no endpoint '%s': %w", endpoint.Name, err)
			}
		}

		if len(endpoint.Headers) > 0 {
			if err := ValidateHeaders(endpoint.Headers); err != nil {
				return nil, fmt.Errorf("erro nos headers do endpoint '%s': %w", endpoint.Name, err)
			}
		}

		if len(endpoint.QueryParams) > 0 {
			if err := ValidateQueryParams(endpoint.QueryParams); err != nil {
				return nil, fmt.Errorf("erro nos query_params do endpoint '%s': %w", endpoint.Name, err)
			}
		}
	}

	if hasProtectedEndpoints && config.Authentication == nil {
		return nil, errors.New("configuração de 'authentication' é obrigatória para endpoints protegidos")
	}

	return &config, nil
}

func ValidateHeaders(headers map[string]string) error {
	validTypes := []string{"string", "int", "float", "bool"}

	for key, typ := range headers {
		if !isValidType(typ, validTypes) {
			return fmt.Errorf("o header '%s' possui um tipo inválido: '%s'", key, typ)
		}
	}

	return nil
}

func ValidateQueryParams(queryParams map[string]string) error {
	validTypes := []string{"string", "int", "float", "bool"}

	for key, typ := range queryParams {
		if !isValidType(typ, validTypes) {
			return fmt.Errorf("o query_param '%s' possui um tipo inválido: '%s'", key, typ)
		}
	}

	return nil
}

func ValidateDatabaseConfig(db DatabaseConfig) error {
	supportedTypes := []string{"sqlite", "mysql", "postgresql"}

	isValidType := false
	for _, typ := range supportedTypes {
		if db.Type == typ {
			isValidType = true
			break
		}
	}

	if !isValidType {
		return fmt.Errorf("tipo de banco de dados '%s' não é suportado. Tipos válidos: %v", db.Type, supportedTypes)
	}

	if db.Type != "sqlite" && (db.Host == "" || db.Port == 0 || db.Username == "" || db.Password == "" || db.Name == "") {
		return errors.New("configurações de 'host', 'port', 'username', 'password' e 'name' são obrigatórias para bancos não SQLite")
	}

	return nil
}

func isValidType(typ string, validTypes []string) bool {
	for _, validType := range validTypes {
		if typ == validType {
			return true
		}
	}
	return false
}

func generateHandlerName(name, method string) string {
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "{", "")
	name = strings.ReplaceAll(name, "}", "")
	return strings.Title(strings.ToLower(method)) + strings.Title(name)
}
