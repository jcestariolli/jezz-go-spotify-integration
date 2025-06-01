package app

import (
	_ "embed"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
	"jezz-go-spotify-integration/internal/auth"
)

const configFileName = "config.yml"

//go:embed config.yml
var configData []byte

const spotifyCliCredentialsFileName = "spotify_client_credentials.yml"

//go:embed spotify_client_credentials.yml
var spotifyCliCredentialsData []byte

type Config struct {
	Clients ClientsConfig `json:"clients" yaml:"clients" validate:"required"`
}

type ClientsConfig struct {
	Spotify SpotifyConfig `json:"spotify" yaml:"spotify" validate:"required"`
}

type SpotifyConfig struct {
	BaseUrl           string `json:"base_url" yaml:"base_url"  validate:"required,url"`
	AccountsUrl       string `json:"accounts_url" yaml:"accounts_url"  validate:"required,url"`
	ClientCredentials auth.ClientCredentials
}

func Load() (Config, error) {
	config, err := loadConfig()
	if err != nil {
		return config, err
	}
	cliCredentials, err2 := loadSpotifyCliCredentials()
	if err2 != nil {
		return config, err2
	}
	config.Clients.Spotify.ClientCredentials = cliCredentials
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
			"error while unmarshalling app config - verify file %s | Error: %w",
			configFileName,
			err,
		)
	}
	return config, nil
}

func loadSpotifyCliCredentials() (auth.ClientCredentials, error) {
	var cliCredentials auth.ClientCredentials
	if err := yaml.Unmarshal(spotifyCliCredentialsData, &cliCredentials); err != nil {
		return cliCredentials, fmt.Errorf(
			"error while unmarshalling spotify client credentials config - verify file %s | Error: %w",
			spotifyCliCredentialsFileName,
			err,
		)
	}
	return cliCredentials, nil
}
