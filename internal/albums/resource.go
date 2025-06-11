package albums

import (
	"fmt"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
	"strings"
)

const (
	apiVersion          = "/v1"
	albumsResource      = "/albums"
	tracksResource      = "/tracks"
	newReleasesResource = "/browse/new-releases"
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
	url := r.baseUrl + apiVersion + albumsResource + "/" + albumId
	queryParams := map[string]string{}
	params := []model.Pair[string, model.StringEvaluator]{
		{"market", market},
	}
	queryParams = utils.AppendQueryParams(queryParams, params...)
	output := &model.Album{}

	if err := utils.DoGetRequestAndValidateSuccess(url, queryParams, accessToken, output); err != nil {
		return model.Album{}, fmt.Errorf("error executing album request for album ID - %s - %w", albumId, err)
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

	url := r.baseUrl + apiVersion + albumsResource
	albumsIdsStr := strings.Join(albumsIds, ",")
	queryParams := map[string]string{
		"ids": albumsIdsStr,
	}
	params := []model.Pair[string, model.StringEvaluator]{
		{"market", market},
	}
	queryParams = utils.AppendQueryParams(queryParams, params...)
	output := &model.MultipleAlbums{}

	if err := utils.DoGetRequestAndValidateSuccess(url, queryParams, accessToken, output); err != nil {
		return []model.Album{}, fmt.Errorf("error executing album request for albums IDs - %s - %w", albumsIdsStr, err)
	}
	return (*output).Albums, nil
}

func (r Resource) GetAlbumTracks(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	limit *model.Limit,
	offset *model.Offset,
	albumId string,
) (model.SimplifiedTracksPaginated, error) {
	if err := utils.ValidatePaginationParams(limit, offset); err != nil {
		return model.SimplifiedTracksPaginated{}, fmt.Errorf("error creating album tracks request for album ID - %s - %w", albumId, err)
	}

	url := r.baseUrl + apiVersion + albumsResource + "/" + albumId + tracksResource
	queryParams := map[string]string{}
	params := []model.Pair[string, model.StringEvaluator]{
		{"market", market},
		{"limit", limit},
		{"offset", offset},
	}
	queryParams = utils.AppendQueryParams(queryParams, params...)
	output := &model.SimplifiedTracksPaginated{}

	if err := utils.DoGetRequestAndValidateSuccess(url, queryParams, accessToken, output); err != nil {
		return model.SimplifiedTracksPaginated{}, fmt.Errorf("error executing album tracks request for album ID - %s - %w", albumId, err)
	}
	return *output, nil
}

func (r Resource) GetNewReleases(
	accessToken model.AccessToken,
	limit *model.Limit,
	offset *model.Offset,
) (model.AlbumsNewRelease, error) {
	if err := utils.ValidatePaginationParams(limit, offset); err != nil {
		return model.AlbumsNewRelease{}, fmt.Errorf("error creating new releases request - %w", err)
	}

	url := r.baseUrl + apiVersion + newReleasesResource
	queryParams := map[string]string{}
	params := []model.Pair[string, model.StringEvaluator]{
		{"limit", limit},
		{"offset", offset},
	}
	queryParams = utils.AppendQueryParams(queryParams, params...)
	output := &model.AlbumsNewRelease{}

	if err := utils.DoGetRequestAndValidateSuccess(url, queryParams, accessToken, output); err != nil {
		return model.AlbumsNewRelease{}, fmt.Errorf("error executing new releases request - %w", err)
	}
	return *output, nil
}

func (r Resource) validateAlbumsIdsLen(albumsIds []string) error {
	if len(albumsIds) < 1 {
		return fmt.Errorf("error getting album - album id must not be null")
	}
	return nil
}
