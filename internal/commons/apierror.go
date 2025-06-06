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
