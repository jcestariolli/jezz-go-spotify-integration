package artists

import (
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"io"
	"jezz-go-spotify-integration/internal/commons"
	"jezz-go-spotify-integration/internal/model"
	"net/http"
	"strings"
)

const (
	apiVersion      = "/v1"
	artistsResource = "/artists"
)

type Resource struct {
	baseUrl string
}

func NewResource(
	baseUrl string,
) Resource {
	return Resource{
		baseUrl: baseUrl,
	}
}

func (a Resource) Get(accessToken model.AccessToken, artistId string) (model.Artist, error) {
	req, cErr := a.createRequest("GET", "/"+artistId, map[string]string{}, accessToken)
	if cErr != nil {
		return model.Artist{}, fmt.Errorf("error creating artist request for astist ID - %s - %w", artistId, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return model.Artist{}, fmt.Errorf("error connecting to artist client for astist ID - %s - %w", artistId, reqErr)
	}

	if vErr := a.validateReqSuccess(resp); vErr != nil {
		return model.Artist{}, vErr
	}
	artist, pErr := a.parseSingleArtistResponse(resp)
	if pErr != nil {
		return model.Artist{}, fmt.Errorf("error parsing response from resource for astist ID - %s - %w", artistId, pErr)
	}
	return artist, nil
}

func (a Resource) GetBatch(accessToken model.AccessToken, artistsIds ...string) (model.Artists, error) {
	if err := a.validateArtistsIdSize(artistsIds); err != nil {
		return model.Artists{}, err
	}
	artistsIdsStr := strings.Join(artistsIds, ",")
	queryParameters := map[string]string{
		"ids": artistsIdsStr,
	}

	req, cErr := a.createRequest("GET", "", queryParameters, accessToken)
	if cErr != nil {
		return model.Artists{}, fmt.Errorf("error creating artist request for astists IDs - %s - %w", artistsIdsStr, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return model.Artists{}, fmt.Errorf("error connecting to artist client for astists IDs - %s - %w", artistsIdsStr, reqErr)
	}

	if vErr := a.validateReqSuccess(resp); vErr != nil {
		return model.Artists{}, vErr
	}

	artists, pErr := a.parseMultipleArtistsResponse(resp)
	if pErr != nil {
		return model.Artists{}, fmt.Errorf("error parsing response from resource  for astists ID - %s - %w", artistsIdsStr, pErr)
	}
	return artists, nil

}

func (a Resource) validateArtistsIdSize(artistIds []string) error {
	if len(artistIds) < 1 {
		return fmt.Errorf("error getting artist - artist id must not be null")
	}
	return nil
}

func (a Resource) createRequest(method string, path string, queryParams map[string]string, accessToken model.AccessToken) (*http.Request, error) {
	url := a.baseUrl + apiVersion + artistsResource
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
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken.String())
	return req, err
}

func (a Resource) validateReqSuccess(resp *http.Response) *commons.ResourceError {
	if resp.StatusCode != 200 {
		apiErr := commons.ResourceError{
			Status:  resp.StatusCode,
			Message: "error in artists API",
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

func (a Resource) parseMultipleArtistsResponse(resp *http.Response) (model.Artists, error) {
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)

	multipleArtists := model.MultipleArtists{}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return multipleArtists.Artists, err
	}

	if err = json.Unmarshal(respBody, &multipleArtists); err != nil {
		var apiErr commons.ResourceError
		if err2 := json.Unmarshal(respBody, &apiErr); err2 == nil && apiErr.Message != "" {
			return multipleArtists.Artists, apiErr
		}
		var errMessage string
		if len(multipleArtists.Artists) < 1 {
			errMessage = ": empty artists list (response body: " + string(respBody) + " )"
		} else {
			errMessage = ", no details were provided"
		}
		return multipleArtists.Artists, commons.AppError{
			Code:    resp.Status,
			Message: "error obtaining artists" + errMessage,
		}
	}
	return multipleArtists.Artists, nil
}

func (a Resource) parseSingleArtistResponse(resp *http.Response) (model.Artist, error) {
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)

	artist := model.Artist{}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return artist, err
	}

	if err = json.Unmarshal(respBody, &artist); err != nil {
		var apiErr commons.ResourceError
		if err2 := json.Unmarshal(respBody, &apiErr); err2 == nil && apiErr.Message != "" {
			return artist, apiErr
		}
		return artist, commons.AppError{
			Code:    resp.Status,
			Message: "error obtaining artists, no details were provided",
		}
	}
	return artist, nil
}
