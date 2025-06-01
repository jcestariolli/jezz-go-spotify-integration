package model

import "encoding/json"

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ApiError) Error() string {
	if body, err := json.Marshal(e); err == nil {
		return string(body)
	}
	return "api error, no details provided"
}
