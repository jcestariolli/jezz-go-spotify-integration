package resource

import (
	"fmt"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/model"
)

type SpotifyTracksResource struct {
	httpClient client.HTTPApiClient
	baseURL    string
}

func NewSpotifyTracksResource(
	httpClient client.HTTPApiClient,
	baseURL string,
) TracksResource {
	return SpotifyTracksResource{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

func (r SpotifyTracksResource) GetTrack(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	trackID model.ID,
) (model.Track, error) {
	url := r.baseURL + APIVersion + TracksPath + "/" + trackID.String()
	queryParams := &model.QueryParams{
		"market": market,
	}
	output := &model.Track{}

	if err := r.httpClient.DoRequest(model.HTTPGet, url, queryParams, client.ContentTypeJSON, &accessToken, output); err != nil {
		return model.Track{}, fmt.Errorf("error executing track request for track ID - %s - %w", trackID.String(), err)
	}
	return *output, nil
}

func (r SpotifyTracksResource) GetTracks(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	tracksIDs model.TracksIDs,
) ([]model.Track, error) {
	if err := r.validateTracksIDsSize(tracksIDs); err != nil {
		return []model.Track{}, err
	}

	url := r.baseURL + APIVersion + TracksPath
	queryParams := &model.QueryParams{
		"ids":    tracksIDs,
		"market": market,
	}
	output := &model.MultipleTracks{}

	if err := r.httpClient.DoRequest(model.HTTPGet, url, queryParams, client.ContentTypeJSON, &accessToken, output); err != nil {
		return []model.Track{}, fmt.Errorf("error executing track request for tracks IDs - %s - %w", tracksIDs.String(), err)
	}
	return output.Tracks, nil
}

func (r SpotifyTracksResource) validateTracksIDsSize(tracksIDs model.TracksIDs) error {
	if len(tracksIDs) < 1 {
		return fmt.Errorf("error getting track - track id must not be null")
	}
	return nil
}
