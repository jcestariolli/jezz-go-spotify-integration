package auth

import (
	"encoding/base64"
	"encoding/json"
)

type ClientCredentials struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// LoadClientCredentialsFromFile reads a JSON config file and returns a ClientCredentials instance
func LoadClientCredentialsFromFile(clientCredentialsConfig []byte) (ClientCredentials, error) {
	var credentials ClientCredentials
	// Unmarshal JSON into ClientCredentials struct
	err := json.Unmarshal(clientCredentialsConfig, &credentials)
	if err != nil {
		return credentials, err
	}

	return credentials, nil
}
func (c ClientCredentials) Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(c.ClientId + ":" + c.ClientSecret))
}

func (c ClientCredentials) GenerateAuthorizationHeader() AuthorizationHeader {
	return AuthorizationHeader{
		Key:   "Authorization",
		Value: BasicAuthorizationType.String() + " " + c.Encode(),
	}
}
