package model

type AccessToken string

func (t AccessToken) String() string {
	return string(t)
}

type Authentication struct {
	AccessToken AccessToken `json:"access_token"`
	TokenType   string      `json:"token_type"`
	ExpiresIn   int         `json:"expires_in"`
}
