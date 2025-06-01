package model

type AccessToken string

type OAuthResponseBody struct {
	AccessToken AccessToken `json:"access_token"`
	TokenType   string      `json:"token_type"`
	ExpiresIn   int         `json:"expires_in"`
}

type OAuthResponse struct {
	Status      string
	StatusCode  int
	SuccessBody *OAuthResponseBody
	ErrorBody   *ErrorResponseBody
}
