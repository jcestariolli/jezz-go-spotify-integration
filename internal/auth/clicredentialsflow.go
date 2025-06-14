package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"jezz-go-spotify-integration/internal/commons"
	"jezz-go-spotify-integration/internal/model"
	"net/http"
	"net/url"
	"strings"
)

const (
	cliCredentialsPath = "/api/token"
)

var (
	// for testing purposes
	httpNewRequest = http.NewRequest
)

type CliCredentialsFlow struct {
	accountURL   string
	clientID     string
	clientSecret string
	httpClient   http.Client
}

func NewCliCredentialsFlow(
	accountURL string,
	clientID string,
	clientSecret string,
) CliCredentialsFlow {
	return CliCredentialsFlow{
		accountURL:   accountURL,
		clientID:     clientID,
		clientSecret: clientSecret,
		httpClient:   http.Client{},
	}
}

func (c CliCredentialsFlow) Authenticate() (*model.Authentication, error) {
	req, err := c.createRequest()
	if err != nil {
		return nil, fmt.Errorf("error creating client credentials request - %w", err)
	}

	resp, err := (&c.httpClient).Do(req)
	if err != nil {
		return nil, fmt.Errorf("error connecting to authorization client - %w", err)
	}

	if err := c.validateRespStatus(resp); err != nil {
		return nil, err
	}

	authResp, err := c.parseResponse(resp)
	if err != nil {
		return authResp, fmt.Errorf("error authenticating - %w", err)
	}
	return authResp, nil
}

func (c CliCredentialsFlow) createRequest() (*http.Request, error) {
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	req, err := httpNewRequest("POST", c.accountURL+cliCredentialsPath, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.clientID, c.clientSecret)
	return req, err
}

func (c CliCredentialsFlow) validateRespStatus(resp *http.Response) error {
	if resp.StatusCode != 200 {
		appErr := commons.AppError{
			Code:    resp.Status,
			Message: "error authenticating",
			Details: "no details were provided",
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return appErr
		}
		var authErr commons.AuthenticationError
		if err = json.Unmarshal(respBody, &authErr); err == nil && authErr.Err != "" {
			appErr.Message = authErr.Err
			appErr.Details = authErr.ErrDescription
		}
		return appErr
	}
	return nil
}

func (c CliCredentialsFlow) parseResponse(resp *http.Response) (*model.Authentication, error) {
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var authBody model.Authentication
	if err = json.Unmarshal(respBody, &authBody); err != nil || authBody.AccessToken == "" {
		return nil, commons.AppError{Code: resp.Status, Message: "error obtaining auth response"}
	}

	return &authBody, nil
}
