package commons

import "encoding/json"

type ResourceError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ResourceError) Error() string {
	if body, err := json.Marshal(e); err == nil {
		return string(body)
	}
	return "resource error, no details provided"
}

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

type AuthenticationError struct {
	Err            string `json:"error"`
	ErrDescription string `json:"error_description"`
}
