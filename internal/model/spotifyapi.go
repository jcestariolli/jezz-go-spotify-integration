package model

type ErrorResponseBody struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type OAuthResponseBody struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type OAuthResponse struct {
	Status      string
	StatusCode  int
	SuccessBody *OAuthResponseBody
	ErrorBody   *ErrorResponseBody
}
