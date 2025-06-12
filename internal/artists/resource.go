package artists

import (
	"fmt"
	"jezz-go-spotify-integration/internal"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/model"
)

type Resource struct {
	httpClient client.HttpClient
	baseUrl    string
}

func NewResource(
	baseUrl string,
) internal.ArtistsResource {
	return Resource{
		httpClient: client.HttpCustomClient{},
		baseUrl:    baseUrl,
	}
}

func (r Resource) GetArtist(
	accessToken model.AccessToken,
	artistId model.Id,
) (model.Artist, error) {
	url := r.baseUrl + internal.ApiVersion + internal.ArtistsPath + "/" + artistId.String()
	output := &model.Artist{}

	if err := r.httpClient.DoRequest(model.HttpGet, url, &model.QueryParams{}, accessToken, output); err != nil {
		return model.Artist{}, fmt.Errorf("error executing artist request for astist ID - %s - %w", artistId.String(), err)
	}
	return *output, nil
}

func (r Resource) GetArtists(
	accessToken model.AccessToken,
	artistsIds model.ArtistsIds,
) ([]model.Artist, error) {
	if err := r.validateArtistsIdSize(artistsIds); err != nil {
		return []model.Artist{}, err
	}

	url := r.baseUrl + internal.ApiVersion + internal.ArtistsPath
	queryParams := &model.QueryParams{
		"ids": artistsIds,
	}
	output := &model.MultipleArtists{}

	if err := r.httpClient.DoRequest(model.HttpGet, url, queryParams, accessToken, output); err != nil {
		return []model.Artist{}, fmt.Errorf("error executing artist request for astists IDs - %s - %w", artistsIds.String(), err)
	}
	return (*output).Artists, nil
}

func (r Resource) GetArtistAlbums(
	accessToken model.AccessToken,
	includeGroups *model.AlbumGroups,
	market *model.AvailableMarket,
	limit *model.Limit,
	offset *model.Offset,
	artistId model.Id,
) (model.SimplifiedArtistAlbumsPaginated, error) {
	url := r.baseUrl + internal.ApiVersion + internal.ArtistsPath + "/" + artistId.String() + internal.AlbumsPath
	queryParams := &model.QueryParams{
		"include_groups": includeGroups,
		"market":         market,
		"limit":          limit,
		"offset":         offset,
	}
	output := &model.SimplifiedArtistAlbumsPaginated{}

	if err := r.httpClient.DoRequest(model.HttpGet, url, queryParams, accessToken, output); err != nil {
		return model.SimplifiedArtistAlbumsPaginated{}, fmt.Errorf("error executing artist albums request for astist ID - %s - %w", artistId.String(), err)
	}
	return *output, nil
}

func (r Resource) GetArtistTopTracks(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	artistId model.Id,
) ([]model.Track, error) {
	url := r.baseUrl + internal.ApiVersion + internal.ArtistsPath + "/" + artistId.String() + internal.TopTracksPath
	queryParams := &model.QueryParams{
		"market": market,
	}
	output := &model.MultipleTracks{}

	if err := r.httpClient.DoRequest(model.HttpGet, url, queryParams, accessToken, output); err != nil {
		return []model.Track{}, fmt.Errorf("error executing artist top-tracks request for astist ID - %s - %w", artistId.String(), err)
	}
	return (*output).Tracks, nil
}

func (r Resource) validateArtistsIdSize(artistsIds model.ArtistsIds) error {
	if len(artistsIds) < 1 {
		return fmt.Errorf("error getting artist - artist id must not be null")
	}
	return nil
}
