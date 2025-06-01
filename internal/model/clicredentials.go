package model

import (
	"encoding/base64"
)

type CliCredentials struct {
	Id     string `json:"client_id" yaml:"client_id" validate:"required"`
	Secret string `json:"client_secret" yaml:"client_secret" validate:"required"`
}

func (c CliCredentials) Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(c.Id + ":" + c.Secret))
}
