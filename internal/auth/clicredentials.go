package auth

import (
	"encoding/base64"
)

type ClientCredentials struct {
	ClientId     string `json:"client_id" yaml:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" yaml:"client_secret" validate:"required"`
}

func (c ClientCredentials) Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(c.ClientId + ":" + c.ClientSecret))
}

func (c ClientCredentials) ToAuthorizationHeader() AuthorizationHeader {
	return AuthorizationHeader{
		Key:   "Authorization",
		Value: BasicAuthorizationType.String() + " " + c.Encode(),
	}
}
