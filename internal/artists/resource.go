package artists

import (
	"fmt"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
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
	req, cErr := utils.CreateHttpRequest(a.baseUrl+apiVersion+artistsResource, "GET", "/"+artistId, map[string]string{}, accessToken)
	if cErr != nil {
		return model.Artist{}, fmt.Errorf("error creating artist request for astist ID - %s - %w", artistId, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return model.Artist{}, fmt.Errorf("error connecting to artist client for astist ID - %s - %w", artistId, reqErr)
	}

	if vErr := utils.ValidateHttpResponseStatus(resp); vErr != nil {
		return model.Artist{}, vErr
	}
	output := &model.Artist{}
	if pErr := utils.ParseHttpResponse(resp, output); pErr != nil {
		return model.Artist{}, fmt.Errorf("error parsing response from resource for astist ID - %s - %w", artistId, pErr)
	}
	return *output, nil
}

func (a Resource) GetBatch(accessToken model.AccessToken, artistsIds ...string) (model.Artists, error) {
	if err := a.validateArtistsIdSize(artistsIds); err != nil {
		return model.Artists{}, err
	}
	artistsIdsStr := strings.Join(artistsIds, ",")
	queryParameters := map[string]string{
		"ids": artistsIdsStr,
	}

	req, cErr := utils.CreateHttpRequest(a.baseUrl+apiVersion+artistsResource, "GET", "", queryParameters, accessToken)
	if cErr != nil {
		return model.Artists{}, fmt.Errorf("error creating artist request for astists IDs - %s - %w", artistsIdsStr, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return model.Artists{}, fmt.Errorf("error connecting to artist client for astists IDs - %s - %w", artistsIdsStr, reqErr)
	}

	if vErr := utils.ValidateHttpResponseStatus(resp); vErr != nil {
		return model.Artists{}, vErr
	}

	output := &model.MultipleArtists{}
	if pErr := utils.ParseHttpResponse(resp, output); pErr != nil {
		return model.Artists{}, fmt.Errorf("error parsing response from resource for astists ID - %s - %w", artistsIdsStr, pErr)
	}
	return output.Artists, nil

}

func (a Resource) validateArtistsIdSize(artistIds []string) error {
	if len(artistIds) < 1 {
		return fmt.Errorf("error getting artist - artist id must not be null")
	}
	return nil
}
