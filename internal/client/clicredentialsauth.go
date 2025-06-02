package client

import (
	"encoding/json"
	"fmt"
	"io"
	"jezz-go-spotify-integration/internal/model"
	"net/http"
	"net/url"
	"strings"
)

const (
	cliCredentialsPath = "/api/token"
)

type CliCredentialsAuthClient struct {
	endpoint                  string
	accountUrl                string
	cliCredentialsEncodedAuth string
	cliCredentials            model.CliCredentials
}

func NewCliCredentialsAuthClient(
	baseUrl string,
	accountUrl string,
	cliCredentials model.CliCredentials,
) CliCredentialsAuthClient {
	return CliCredentialsAuthClient{
		endpoint:       baseUrl,
		accountUrl:     accountUrl,
		cliCredentials: cliCredentials,
	}
}

func (c CliCredentialsAuthClient) Authenticate() (*model.AuthSession, error) {
	authSession := &model.AuthSession{}
	req, err := c.createRequest()
	if err != nil {
		return authSession, fmt.Errorf("error creating client credentials request - %w", err)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return authSession, fmt.Errorf("error connecting to authorization client - %w", err)
	}

	if err := c.validateRespStatus(resp); err != nil {
		return authSession, err
	}

	authSession, err = c.parseResponse(resp)
	if err != nil {
		return authSession, fmt.Errorf("error authenticating - %w", err)
	}
	return authSession, nil
}

func (c CliCredentialsAuthClient) createRequest() (*http.Request, error) {
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", c.accountUrl+cliCredentialsPath, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.cliCredentials.Id, c.cliCredentials.Secret)
	return req, err
}

func (c CliCredentialsAuthClient) validateRespStatus(resp *http.Response) error {
	if resp.StatusCode != 200 {
		appErr := model.AppError{
			Code:    resp.Status,
			Message: "error authenticating",
			Details: "no details were provided",
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return appErr
		}
		var authErr model.AuthError
		if err = json.Unmarshal(respBody, &authErr); err == nil && authErr.Err != "" {
			appErr.Message = authErr.Err
			appErr.Details = authErr.ErrDescription
		}
		return appErr
	}
	return nil
}

func (c CliCredentialsAuthClient) parseResponse(resp *http.Response) (*model.AuthSession, error) {
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)

	var authSession *model.AuthSession
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return authSession, err
	}
	var authBody model.Auth
	if err = json.Unmarshal(respBody, &authBody); err != nil || authBody.AccessToken == "" {
		return authSession, model.AppError{Code: resp.Status, Message: "error obtaining auth response"}
	}
	authSession = &model.AuthSession{}
	authSession.Auth = &authBody

	return authSession, nil
}
