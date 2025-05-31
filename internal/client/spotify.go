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

type SpotifyClient struct {
	url                                  string
	clientCredentialsAuthorizationHeader auth.AuthorizationHeader
}

func NewSpotifyClient(
	apiUrl string,
	clientCredentials auth.ClientCredentials,
) SpotifyClient {
	return SpotifyClient{
		url:                                  apiUrl,
		clientCredentialsAuthorizationHeader: clientCredentials.GenerateAuthorizationHeader(),
	}
}

func (c *SpotifyClient) AuthenticateWithClientCredentials() (auth.AccessToken, error) {
	spotifyResponse, err := c.getClientCredentialsToken()
	if err != nil {
		fmt.Println("Error authenticating")
		return "", err
	}
	return auth.AccessToken(spotifyResponse.AccessToken), err
}

func (c *SpotifyClient) getClientCredentialsToken() (model.ClientCredentialsResponse, error) {

	req, err := c.createClientCredentialsRequest()
	if err != nil {
		return model.ClientCredentialsResponse{}, err
	}

	// Create HTTP c and do the request
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return model.ClientCredentialsResponse{}, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		err = errors.New("Status Code: " + strconv.Itoa(resp.StatusCode) + " | Error: " + string(body))
		fmt.Println("Error making HTTP request:", err)
		return model.ClientCredentialsResponse{}, err
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return model.ClientCredentialsResponse{}, err
	}

	// Process response (example: print it)
	fmt.Println("ClientCredentialsResponse:", string(body))

	spotifyResponse := &model.ClientCredentialsResponse{}
	err = json.Unmarshal(body, spotifyResponse)

	return *spotifyResponse, nil
}

func (c *SpotifyClient) createClientCredentialsRequest() (*http.Request, error) {
	// Prepare form data or JSON payload
	formData := url.Values{}
	formData.Set("grant_type", "client_credentials")

	// Create HTTP request
	req, err := http.NewRequest("POST", c.url, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(c.clientCredentialsAuthorizationHeader.Key, c.clientCredentialsAuthorizationHeader.Value)
	return req, err
}
