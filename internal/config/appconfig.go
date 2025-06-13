package config

type AppConfig struct {
	Client CliConfig `json:"client" yaml:"client" validate:"required"`
}
type CliConfig struct {
	BaseUrl     string `json:"base_url" yaml:"base_url"  validate:"required,url"`
	AccountsUrl string `json:"accounts_url" yaml:"accounts_url"  validate:"required,url"`
}

type AppConfigLoader struct{}

func (a AppConfigLoader) Load(appConfigData []byte) (AppConfig, error) {
	config := AppConfig{}
	if err := loadConfig(appConfigData, &config); err != nil {
		return AppConfig{}, err
	}
	if err := validate(config); err != nil {
		return AppConfig{}, err
	}
	return config, nil
}
