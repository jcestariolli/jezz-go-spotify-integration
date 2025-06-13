package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type Loader[T any] interface {
	Load(configData []byte) (T, error)
}

func validate(t any) error {
	return validator.New().Struct(t)
}

func loadConfig[T any](configData []byte, config *T) error {
	if err := yaml.Unmarshal(configData, config); err != nil {
		return fmt.Errorf("error while unmarshalling config - %w", err)
	}
	return nil
}
