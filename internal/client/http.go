package client

import (
	"encoding/json"
	"fmt"
	"io"
	"jezz-go-spotify-integration/internal/commons"
	"jezz-go-spotify-integration/internal/model"
	"net/http"
	"reflect"
	"strings"

	"github.com/samber/lo"
)

type HTTPClient interface {
	DoRequest(
		method model.HTTPMethod,
		url string,
		queryParams *model.QueryParams,
		accessToken model.AccessToken,
		responseTypedOutput any,
	) error
}

type HTTPCustomClient struct{}

func (c HTTPCustomClient) DoRequest(
	method model.HTTPMethod,
	url string,
	queryParams *model.QueryParams,
	accessToken model.AccessToken,
	responseTypedOutput any,
) error {
	req, cErr := c.createRequest(method, url, queryParams, accessToken)
	if cErr != nil {
		return fmt.Errorf("error creating request - %s", cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return fmt.Errorf("error executing request - %w", reqErr)
	}

	if vErr := c.validateResponseStatus(resp); vErr != nil {
		return vErr
	}

	if pErr := c.parseResponse(resp, responseTypedOutput); pErr != nil {
		return fmt.Errorf("error parsing response - %w", pErr)
	}
	return nil
}

func (c HTTPCustomClient) createRequest(
	method model.HTTPMethod,
	url string,
	queryParams *model.QueryParams,
	accessToken model.AccessToken,
) (*http.Request, error) {
	queryParamsMap := parseQueryParams(queryParams)
	if len(queryParamsMap) > 0 {
		url += "?" + strings.Join(
			lo.MapToSlice(queryParamsMap, func(key string, value string) string {
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

func (c HTTPCustomClient) validateResponseStatus(resp *http.Response) *commons.ResourceError {
	if resp.StatusCode >= 300 {
		apiErr := commons.ResourceError{
			Status:  resp.StatusCode,
			Message: "API http status is not success",
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

func (c HTTPCustomClient) parseResponse(resp *http.Response, output any) error {
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

func parseQueryParams(queryParams *model.QueryParams) map[string]string {
	queryParamsMap := map[string]string{}
	if queryParams != nil {
		for key, stringEvaluator := range *queryParams {
			if stringEvaluator != nil {
				val := reflect.ValueOf(stringEvaluator)
				// Check if the interface holds a nil pointer
				if val.Kind() == reflect.Ptr && val.IsNil() {
					continue
				}
				queryParamsMap[key] = stringEvaluator.String()
			}
		}
	}
	return queryParamsMap
}
