package client

import (
	"encoding/json"
	"fmt"
	"io"
	"jezz-go-spotify-integration/internal/model"
	"net/http"
)

type ArtistsAPIClient struct {
	baseUrl string
}

func NewArtistsAPIClient(
	baseUrl string,
) ArtistsAPIClient {
	return ArtistsAPIClient{
		baseUrl: baseUrl + "/v1/artists",
	}
}

func (a ArtistsAPIClient) GetArtist(accessToken model.AccessToken, artistId string) (model.Artist, error) {
	artist := model.Artist{}
	req, err := a.genArtistRequest(accessToken, artistId)
	if err != nil {
		return artist, fmt.Errorf("error creating artist request - %w", err)
	}

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return artist, fmt.Errorf("error connecting to artist client - %w", err)
	}

	if err := a.validateReqSuccess(resp); err != nil {
		return artist, err
	}

	artist, err = a.parseResponse(resp)
	if err != nil {
		return artist, fmt.Errorf("error for artist id %s - %w", artistId, err)
	}

	return artist, nil
}

func (a ArtistsAPIClient) genArtistRequest(accessToken model.AccessToken, artistId string) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", a.baseUrl, artistId)
	fmt.Println("URL: " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken.String())
	return req, err
}

func (a ArtistsAPIClient) validateReqSuccess(resp *http.Response) error {
	if resp.StatusCode != 200 {
		appErr := model.AppError{
			Code:    resp.Status,
			Message: "error getting artist",
			Details: "no details were provided",
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return appErr
		}
		var apiErr model.ApiError
		if err2 := json.Unmarshal(respBody, &apiErr); err2 == nil && apiErr.Message != "" {
			appErr.Message = apiErr.Message
			appErr.Details = apiErr.Message

		}
		return appErr
	}
	return nil
}

func (a ArtistsAPIClient) parseResponse(resp *http.Response) (model.Artist, error) {
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)

	artist := model.Artist{}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return artist, err
	}

	if err = json.Unmarshal(respBody, &artist); err != nil || artist.Id == "" {
		var apiErr model.ApiError
		if err2 := json.Unmarshal(respBody, &apiErr); err2 == nil && apiErr.Message != "" {
			return artist, apiErr
		}
		return artist, model.AppError{
			Code:    resp.Status,
			Message: "error obtaining artist, no details were provided",
		}
	}
	return artist, nil
}
