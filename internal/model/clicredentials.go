package model

type CliCredentials struct {
	Id     string `json:"client_id" yaml:"client_id" validate:"required"`
	Secret string `json:"client_secret" yaml:"client_secret" validate:"required"`
}
