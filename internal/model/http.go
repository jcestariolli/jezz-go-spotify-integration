package model

const (
	HttpGet HttpMethod = "GET"
)

type HttpMethod string

func (m HttpMethod) String() string {
	return string(m)
}

type QueryParams map[string]StringEvaluator
