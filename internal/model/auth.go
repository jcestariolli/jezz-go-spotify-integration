package model

type AccessToken string

type Auth struct {
	AccessToken AccessToken `json:"access_token"`
	TokenType   string      `json:"token_type"`
	ExpiresIn   int         `json:"expires_in"`
}

type AuthSession struct {
	Auth *Auth
}

func (t AccessToken) String() string {
	return string(t)
}
