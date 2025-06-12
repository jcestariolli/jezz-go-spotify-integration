package config

import (
	_ "embed"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

const configFileName = "config.yml"

//go:embed config.yml
var configData []byte

const spotifyCliCredentialsFileName = "spotify_client_credentials.yml"

//go:embed spotify_client_credentials.yml
var spotifyCliCredentialsData []byte

type Config struct {
	Client CliConfig `json:"client" yaml:"client" validate:"required"`
}
type CliConfig struct {
	BaseUrl        string `json:"base_url" yaml:"base_url"  validate:"required,url"`
	AccountsUrl    string `json:"accounts_url" yaml:"accounts_url"  validate:"required,url"`
	CliCredentials CliCredentials
}

func Load() (Config, error) {
	config, err := loadConfig()
	if err != nil {
		return config, err
	}
	cliCredentials, err := loadSpotifyCliCredentials()
	if err != nil {
		return config, err
	}
	config.Client.CliCredentials = cliCredentials
	err = config.Validate()
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func loadConfig() (Config, error) {
	var config Config
	if err := yaml.Unmarshal(configData, &config); err != nil {
		return config, fmt.Errorf(
			"error while unmarshalling config config - verify file %s | Error: %w",
			configFileName,
			err,
		)
	}
	return config, nil
}

func loadSpotifyCliCredentials() (CliCredentials, error) {
	var cliCredentials CliCredentials
	if err := yaml.Unmarshal(spotifyCliCredentialsData, &cliCredentials); err != nil {
		return cliCredentials, fmt.Errorf(
			"error while unmarshalling spotify client credentials config - verify file %s | Error: %w",
			spotifyCliCredentialsFileName,
			err,
		)
	}
	return cliCredentials, nil
}
