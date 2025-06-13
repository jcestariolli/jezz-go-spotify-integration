package model

const (
	HTTPGet HTTPMethod = "GET"
)

type HTTPMethod string

func (m HTTPMethod) String() string {
	return string(m)
}

type QueryParams map[string]StringEvaluator
