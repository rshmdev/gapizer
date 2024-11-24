package validator

import (
	"errors"
	"strings"

	"gapizer/internal/parser"
)

func ValidateConfig(config *parser.Config) error {
	if config.AppName == "" || config.Port == 0 {
		return errors.New("app_name e port são obrigatórios")
	}

	if len(config.Endpoints) == 0 {
		return errors.New("pelo menos um endpoint deve ser definido")
	}

	for _, endpoint := range config.Endpoints {
		if endpoint.Name == "" || endpoint.Method == "" {
			return errors.New("cada endpoint deve ter nome e método definidos")
		}
		if !strings.HasPrefix(endpoint.Name, "/") {
			return errors.New("o nome do endpoint deve começar com '/'")
		}
	}

	return nil
}
