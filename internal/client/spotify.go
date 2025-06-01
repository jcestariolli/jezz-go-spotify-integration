package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/model"
	"net/http"
	"net/url"
	"strings"
)

const (
	authTokenPath = "/api/token"
)

type SpotifyClient struct {
	baseUrl                  string
	accountUrl               string
	cliCredentialsAuthHeader auth.AuthorizationHeader
}

func NewSpotifyClient(
	baseUrl string,
	accountUrl string,
	clientCredentials auth.ClientCredentials,
) SpotifyClient {
	return SpotifyClient{
		baseUrl:                  baseUrl,
		accountUrl:               accountUrl,
		cliCredentialsAuthHeader: clientCredentials.ToAuthorizationHeader(),
	}
}

func (c *SpotifyClient) AuthenticateWithClientCredentials() (*model.OAuthResponse, error) {
	req, err := c.getCliCredentialsOAuthRequest()
	if err != nil {
		return nil, fmt.Errorf("error while creating client credentials request: %w", err)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while authenticating with client credentials: %w", err)
	}

	oAuthResp, err := c.parseOAuthResponse(resp, err)
	if err != nil {
		return nil, fmt.Errorf("error while parsing client credentials response: %w", err)
	}

	if err := c.validateOAuthResponse(oAuthResp); err != nil {
		return nil, fmt.Errorf("authentication with client credentials failed: %w", err)
	}

	return oAuthResp, nil
}

func (c *SpotifyClient) validateOAuthResponse(oAuthResp *model.OAuthResponse) error {
	if oAuthResp.StatusCode != 200 {
		if errMsg, err := json.Marshal(oAuthResp); err != nil {
			return errors.New("status " + oAuthResp.Status)
		} else {
			return errors.New(string(errMsg))
		}
	}
	return nil
}

func (c *SpotifyClient) parseOAuthResponse(resp *http.Response, err error) (*model.OAuthResponse, error) {
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)
	oAuthResponse := &model.OAuthResponse{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
	}
	var oAuthBody model.OAuthResponseBody
	var errorBody model.ErrorResponseBody
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err4 := json.Unmarshal(respBody, &oAuthBody); err4 != nil || oAuthBody.AccessToken == "" {
		oAuthResponse.SuccessBody = nil
		if err2 := json.Unmarshal(respBody, &errorBody); err2 != nil || errorBody.Error == "" {
			oAuthResponse.ErrorBody = nil
		} else {
			oAuthResponse.ErrorBody = &errorBody
		}
	} else {
		oAuthResponse.SuccessBody = &oAuthBody
	}
	return oAuthResponse, nil
}

func (c *SpotifyClient) getCliCredentialsOAuthRequest() (*http.Request, error) {
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", c.accountUrl+authTokenPath, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(c.cliCredentialsAuthHeader.Key, c.cliCredentialsAuthHeader.Value)
	return req, err
}
