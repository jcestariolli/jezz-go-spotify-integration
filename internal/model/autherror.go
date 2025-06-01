package model

import "encoding/json"

type AuthError struct {
	Err            string `json:"error"`
	ErrDescription string `json:"error_description"`
}

func (e AuthError) Error() string {
	if body, err := json.Marshal(e); err == nil {
		return string(body)
	}
	return "authentication error, no details provided"
}
