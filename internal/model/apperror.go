package model

import "encoding/json"

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func (e AppError) Error() string {
	if body, err := json.Marshal(e); err == nil {
		return string(body)
	}
	return "app error, no details provided"
}
