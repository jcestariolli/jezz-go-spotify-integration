package utils

import (
	"encoding/json"
	"github.com/samber/lo"
	"io"
	"jezz-go-spotify-integration/internal/commons"
	"jezz-go-spotify-integration/internal/model"
	"net/http"
	"reflect"
	"strings"
)

type HttpMethod string

const (
	HttpGet HttpMethod = "GET"
)

func (m HttpMethod) String() string {
	return string(m)
}

func CreateHttpRequest(
	method HttpMethod,
	url string,
	path string,
	queryParams map[string]string,
	accessToken model.AccessToken,
) (*http.Request, error) {
	if path != "" {
		url += path
	}
	if len(queryParams) > 0 {
		url += "?" + strings.Join(
			lo.MapToSlice(queryParams, func(key string, value string) string {
				return key + "=" + value
			}),
			"&",
		)
	}
	req, err := http.NewRequest(method.String(), url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken.String())
	return req, err
}

func ValidateHttpResponseStatus(resp *http.Response) *commons.ResourceError {
	if resp.StatusCode != 200 {
		apiErr := commons.ResourceError{
			Status:  resp.StatusCode,
			Message: "API http status is not 200",
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return &apiErr
		}

		_ = json.Unmarshal(respBody, &apiErr)
		return &apiErr
	}
	return nil
}

func ParseHttpResponse[T any](resp *http.Response, output *T) error {
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(respBody, output); err != nil {
		var apiErr commons.ResourceError
		if err2 := json.Unmarshal(respBody, &apiErr); err2 == nil && apiErr.Message != "" {
			return apiErr
		}
		return commons.AppError{
			Code:    resp.Status,
			Message: "error parsing http response, no details were provided",
		}
	}
	return nil
}

func AppendQueryParams(queryParams map[string]string, stringParams ...model.Pair[string, model.StringEvaluator]) map[string]string {
	for _, pair := range stringParams {
		if pair.Value != nil {
			val := reflect.ValueOf(pair.Value)
			// Check if the interface holds a nil pointer
			if val.Kind() == reflect.Ptr && val.IsNil() {
				continue
			}
			queryParams[pair.Key] = pair.Value.String()
		}
	}
	return queryParams
}
