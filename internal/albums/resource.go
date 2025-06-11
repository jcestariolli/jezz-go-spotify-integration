package albums

import (
	"fmt"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
	"net/http"
	"strings"
)

const (
	apiVersion     = "/v1"
	albumsResource = "/albums"
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

func (r Resource) GetAlbum(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	albumId string,
) (model.Album, error) {
	queryParameters := map[string]string{}
	if market != nil {
		queryParameters["market"] = (*market).String()
	}
	req, cErr := utils.CreateHttpRequest(utils.HttpGet, r.baseUrl+apiVersion+albumsResource, "/"+albumId, queryParameters, accessToken)
	if cErr != nil {
		return model.Album{}, fmt.Errorf("error creating album request for album ID - %s - %w", albumId, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return model.Album{}, fmt.Errorf("error connecting to album client for album ID - %s - %w", albumId, reqErr)
	}

	if vErr := utils.ValidateHttpResponseStatus(resp); vErr != nil {
		return model.Album{}, vErr
	}
	output := &model.Album{}
	if pErr := utils.ParseHttpResponse(resp, output); pErr != nil {
		return model.Album{}, fmt.Errorf("error parsing response from resource for album ID - %s - %w", albumId, pErr)
	}
	return *output, nil
}

func (r Resource) GetAlbums(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	albumsIds ...string,
) ([]model.Album, error) {
	if err := r.validateAlbumsIdsLen(albumsIds); err != nil {
		return []model.Album{}, err
	}

	albumsIdsStr := strings.Join(albumsIds, ",")
	queryParameters := map[string]string{
		"ids": albumsIdsStr,
	}
	if market != nil {
		queryParameters["market"] = (*market).String()
	}
	req, cErr := utils.CreateHttpRequest(utils.HttpGet, r.baseUrl+apiVersion+albumsResource, "", queryParameters, accessToken)
	if cErr != nil {
		return []model.Album{}, fmt.Errorf("error creating album request for albums IDs - %s - %w", albumsIdsStr, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return []model.Album{}, fmt.Errorf("error connecting to album client for albums IDs - %s - %w", albumsIdsStr, reqErr)
	}

	if vErr := utils.ValidateHttpResponseStatus(resp); vErr != nil {
		return []model.Album{}, vErr
	}

	output := &model.MultipleAlbums{}
	if pErr := utils.ParseHttpResponse(resp, output); pErr != nil {
		return []model.Album{}, fmt.Errorf("error parsing response from resource for albums ID - %s - %w", albumsIdsStr, pErr)
	}
	return output.Albums, nil

}

func (r Resource) validateAlbumsIdsLen(albumsIds []string) error {
	if len(albumsIds) < 1 {
		return fmt.Errorf("error getting album - album id must not be null")
	}
	return nil
}
