package albums

import (
	"fmt"
	"jezz-go-spotify-integration/internal"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
)

type Resource struct {
	httpClient client.HttpClient
	baseUrl    string
}

func NewResource(
	baseUrl string,
) internal.AlbumsResource {
	return Resource{
		httpClient: client.HttpCustomClient{},
		baseUrl:    baseUrl,
	}
}

func (r Resource) GetAlbum(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	albumId model.Id,
) (model.Album, error) {
	url := r.baseUrl + internal.ApiVersion + internal.AlbumsPath + "/" + albumId.String()
	queryParams := &model.QueryParams{
		"market": market,
	}
	output := &model.Album{}

	if err := r.httpClient.DoRequest(model.HttpGet, url, queryParams, accessToken, output); err != nil {
		return model.Album{}, fmt.Errorf("error executing album request for album ID - %s - %w", albumId.String(), err)
	}
	return *output, nil
}

func (r Resource) GetAlbums(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	albumsIds model.AlbumsIds,
) ([]model.Album, error) {
	if err := r.validateAlbumsIdsLen(albumsIds); err != nil {
		return []model.Album{}, err
	}

	url := r.baseUrl + internal.ApiVersion + internal.AlbumsPath
	queryParams := &model.QueryParams{
		"ids":    albumsIds,
		"market": market,
	}
	output := &model.MultipleAlbums{}

	if err := r.httpClient.DoRequest(model.HttpGet, url, queryParams, accessToken, output); err != nil {
		return []model.Album{}, fmt.Errorf("error executing album request for albums IDs - %s - %w", albumsIds.String(), err)
	}
	return (*output).Albums, nil
}

func (r Resource) GetAlbumTracks(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	limit *model.Limit,
	offset *model.Offset,
	albumId model.Id,
) (model.SimplifiedTracksPaginated, error) {
	if err := utils.ValidatePaginationParams(limit, offset); err != nil {
		return model.SimplifiedTracksPaginated{}, fmt.Errorf("error creating album tracks request for album ID - %s - %w", albumId.String(), err)
	}

	url := r.baseUrl + internal.ApiVersion + internal.AlbumsPath + "/" + albumId.String() + internal.TracksPath
	queryParams := &model.QueryParams{
		"market": market,
		"limit":  limit,
		"offset": offset,
	}
	output := &model.SimplifiedTracksPaginated{}

	if err := r.httpClient.DoRequest(model.HttpGet, url, queryParams, accessToken, output); err != nil {
		return model.SimplifiedTracksPaginated{}, fmt.Errorf("error executing album tracks request for album ID - %s - %w", albumId.String(), err)
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

	url := r.baseUrl + internal.ApiVersion + internal.NewReleasesPath
	queryParams := &model.QueryParams{
		"limit":  limit,
		"offset": offset,
	}
	output := &model.AlbumsNewRelease{}

	if err := r.httpClient.DoRequest(model.HttpGet, url, queryParams, accessToken, output); err != nil {
		return model.AlbumsNewRelease{}, fmt.Errorf("error executing new releases request - %w", err)
	}
	return *output, nil
}

func (r Resource) validateAlbumsIdsLen(albumsIds model.AlbumsIds) error {
	if len(albumsIds) < 1 {
		return fmt.Errorf("error getting album - album id must not be null")
	}
	return nil
}
