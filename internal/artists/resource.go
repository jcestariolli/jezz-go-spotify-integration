package artists

import (
	"fmt"
	"jezz-go-spotify-integration/internal"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/model"
)

type Resource struct {
	httpClient client.HTTPClient
	baseURL    string
}

func NewResource(
	baseURL string,
) internal.ArtistsResource {
	return Resource{
		httpClient: client.HTTPCustomClient{},
		baseURL:    baseURL,
	}
}

func (r Resource) GetArtist(
	accessToken model.AccessToken,
	artistID model.ID,
) (model.Artist, error) {
	url := r.baseURL + internal.APIVersion + internal.ArtistsPath + "/" + artistID.String()
	output := &model.Artist{}

	if err := r.httpClient.DoRequest(model.HTTPGet, url, &model.QueryParams{}, accessToken, output); err != nil {
		return model.Artist{}, fmt.Errorf("error executing artist request for astist ID - %s - %w", artistID.String(), err)
	}
	return *output, nil
}

func (r Resource) GetArtists(
	accessToken model.AccessToken,
	artistsIDs model.ArtistsIDs,
) ([]model.Artist, error) {
	if err := r.validateArtistsIDsSize(artistsIDs); err != nil {
		return []model.Artist{}, err
	}

	url := r.baseURL + internal.APIVersion + internal.ArtistsPath
	queryParams := &model.QueryParams{
		"ids": artistsIDs,
	}
	output := &model.MultipleArtists{}

	if err := r.httpClient.DoRequest(model.HTTPGet, url, queryParams, accessToken, output); err != nil {
		return []model.Artist{}, fmt.Errorf("error executing artist request for astists IDs - %s - %w", artistsIDs.String(), err)
	}
	return output.Artists, nil
}

func (r Resource) GetArtistAlbums(
	accessToken model.AccessToken,
	includeGroups *model.AlbumGroups,
	market *model.AvailableMarket,
	limit *model.Limit,
	offset *model.Offset,
	artistID model.ID,
) (model.SimplifiedArtistAlbumsPaginated, error) {
	url := r.baseURL + internal.APIVersion + internal.ArtistsPath + "/" + artistID.String() + internal.AlbumsPath
	queryParams := &model.QueryParams{
		"include_groups": includeGroups,
		"market":         market,
		"limit":          limit,
		"offset":         offset,
	}
	output := &model.SimplifiedArtistAlbumsPaginated{}

	if err := r.httpClient.DoRequest(model.HTTPGet, url, queryParams, accessToken, output); err != nil {
		return model.SimplifiedArtistAlbumsPaginated{}, fmt.Errorf("error executing artist albums request for astist ID - %s - %w", artistID.String(), err)
	}
	return *output, nil
}

func (r Resource) GetArtistTopTracks(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	artistID model.ID,
) ([]model.Track, error) {
	url := r.baseURL + internal.APIVersion + internal.ArtistsPath + "/" + artistID.String() + internal.TopTracksPath
	queryParams := &model.QueryParams{
		"market": market,
	}
	output := &model.MultipleTracks{}

	if err := r.httpClient.DoRequest(model.HTTPGet, url, queryParams, accessToken, output); err != nil {
		return []model.Track{}, fmt.Errorf("error executing artist top-tracks request for astist ID - %s - %w", artistID.String(), err)
	}
	return output.Tracks, nil
}

func (r Resource) validateArtistsIDsSize(artistsIDs model.ArtistsIDs) error {
	if len(artistsIDs) < 1 {
		return fmt.Errorf("error getting artist - artist id must not be null")
	}
	return nil
}
