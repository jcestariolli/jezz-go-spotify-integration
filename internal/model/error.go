package model

type ErrorResponseBody struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
