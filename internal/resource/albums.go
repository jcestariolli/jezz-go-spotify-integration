package resource

import (
	"fmt"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
)

type SpotifyAlbumsResource struct {
	httpClient client.HTTPClient
	baseURL    string
}

func NewSpotifyAlbumsResource(
	baseURL string,
) AlbumsResource {
	return SpotifyAlbumsResource{
		httpClient: client.HTTPCustomClient{},
		baseURL:    baseURL,
	}
}

func (r SpotifyAlbumsResource) GetAlbum(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	albumID model.ID,
) (model.Album, error) {
	url := r.baseURL + APIVersion + AlbumsPath + "/" + albumID.String()
	queryParams := &model.QueryParams{
		"market": market,
	}
	output := &model.Album{}

	if err := r.httpClient.DoRequest(model.HTTPGet, url, queryParams, accessToken, output); err != nil {
		return model.Album{}, fmt.Errorf("error executing album request for album ID - %s - %w", albumID.String(), err)
	}
	return *output, nil
}

func (r SpotifyAlbumsResource) GetAlbums(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	albumsIDs model.AlbumsIDs,
) ([]model.Album, error) {
	if err := r.validateAlbumsIDsLen(albumsIDs); err != nil {
		return []model.Album{}, err
	}

	url := r.baseURL + APIVersion + AlbumsPath
	queryParams := &model.QueryParams{
		"ids":    albumsIDs,
		"market": market,
	}
	output := &model.MultipleAlbums{}

	if err := r.httpClient.DoRequest(model.HTTPGet, url, queryParams, accessToken, output); err != nil {
		return []model.Album{}, fmt.Errorf("error executing album request for albums IDs - %s - %w", albumsIDs.String(), err)
	}
	return output.Albums, nil
}

func (r SpotifyAlbumsResource) GetAlbumTracks(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	limit *model.Limit,
	offset *model.Offset,
	albumID model.ID,
) (model.SimplifiedTracksPaginated, error) {
	if err := utils.ValidatePaginationParams(limit, offset); err != nil {
		return model.SimplifiedTracksPaginated{}, fmt.Errorf("error creating album tracks request for album ID - %s - %w", albumID.String(), err)
	}

	url := r.baseURL + APIVersion + AlbumsPath + "/" + albumID.String() + TracksPath
	queryParams := &model.QueryParams{
		"market": market,
		"limit":  limit,
		"offset": offset,
	}
	output := &model.SimplifiedTracksPaginated{}

	if err := r.httpClient.DoRequest(model.HTTPGet, url, queryParams, accessToken, output); err != nil {
		return model.SimplifiedTracksPaginated{}, fmt.Errorf("error executing album tracks request for album ID - %s - %w", albumID.String(), err)
	}
	return *output, nil
}

func (r SpotifyAlbumsResource) GetNewReleases(
	accessToken model.AccessToken,
	limit *model.Limit,
	offset *model.Offset,
) (model.AlbumsNewRelease, error) {
	if err := utils.ValidatePaginationParams(limit, offset); err != nil {
		return model.AlbumsNewRelease{}, fmt.Errorf("error creating new releases request - %w", err)
	}

	url := r.baseURL + APIVersion + NewReleasesPath
	queryParams := &model.QueryParams{
		"limit":  limit,
		"offset": offset,
	}
	output := &model.AlbumsNewRelease{}

	if err := r.httpClient.DoRequest(model.HTTPGet, url, queryParams, accessToken, output); err != nil {
		return model.AlbumsNewRelease{}, fmt.Errorf("error executing new releases request - %w", err)
	}
	return *output, nil
}

func (r SpotifyAlbumsResource) validateAlbumsIDsLen(albumsIDs model.AlbumsIDs) error {
	if len(albumsIDs) < 1 {
		return fmt.Errorf("error getting album - album id must not be null")
	}
	return nil
}
