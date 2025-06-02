package client

import (
	"encoding/json"
	"fmt"
	"io"
	"jezz-go-spotify-integration/internal/model"
	"net/http"
)

type ArtistsResourceCli struct {
	baseUrl string
}

func NewArtistsResourceCli(
	baseUrl string,
) ArtistsResourceCli {
	return ArtistsResourceCli{
		baseUrl: baseUrl + "/v1/artists",
	}
}

func (a ArtistsResourceCli) GetArtist(accessToken model.AccessToken, artistId string) (model.Artist, error) {
	artist := model.Artist{}
	req, err := a.createRequest("GET", "/"+artistId, accessToken)
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

func (a ArtistsResourceCli) createRequest(method string, path string, accessToken model.AccessToken) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", a.baseUrl, path)
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken.String())
	return req, err
}

func (a ArtistsResourceCli) validateReqSuccess(resp *http.Response) *model.ApiError {
	if resp.StatusCode != 200 {
		apiErr := model.ApiError{
			Status:  resp.StatusCode,
			Message: "error getting artist",
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

func (a ArtistsResourceCli) parseResponse(resp *http.Response) (model.Artist, error) {
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
