package config

type CliCredentials struct {
	Id     string `json:"client_id" yaml:"client_id" validate:"required"`
	Secret string `json:"client_secret" yaml:"client_secret" validate:"required"`
}

type CliCredentialsConfigLoader struct{}

func (a CliCredentialsConfigLoader) Load(cliCredConfigData []byte) (CliCredentials, error) {
	config := CliCredentials{}
	if err := loadConfig(cliCredConfigData, &config); err != nil {
		return CliCredentials{}, err
	}
	if err := validate(config); err != nil {
		return CliCredentials{}, err
	}
	return config, nil
}
