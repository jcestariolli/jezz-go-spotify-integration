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
	"strconv"
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

func (c *SpotifyClient) AuthenticateWithClientCredentials() (model.ClientCredentialsResponse, error) {
	var response model.ClientCredentialsResponse
	req, err := c.createClientCredentialsRequest()
	if err != nil {
		return response, fmt.Errorf("error while creating client credentials request: %w", err)
	}

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return response, fmt.Errorf("error while authenticating with client credentials: %w", err)
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		err = errors.New("authentication failed - Status Code: " + strconv.Itoa(resp.StatusCode) + " | Body: " + string(body))
		return model.ClientCredentialsResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("error reading response body: %w", err)
	}

	if json.Unmarshal(body, &response) != nil {
		return response, fmt.Errorf("error unmarshalling response body to model: %w", err)
	}

	return *&response, nil
}

func (c *SpotifyClient) createClientCredentialsRequest() (*http.Request, error) {
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
