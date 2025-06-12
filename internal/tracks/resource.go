package tracks

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
) Resource {
	return Resource{
		httpClient: client.HttpCustomClient{},
		baseUrl:    baseUrl,
	}
}

func (r Resource) GetTrack(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	trackId model.Id,
) (model.Track, error) {
	url := r.baseUrl + internal.ApiVersion + internal.TracksPath + "/" + trackId.String()
	queryParams := &model.QueryParams{
		"market": market,
	}
	output := &model.Track{}

	if err := r.httpClient.DoRequest(model.HttpGet, url, queryParams, accessToken, output); err != nil {
		return model.Track{}, fmt.Errorf("error executing track request for track ID - %s - %w", trackId.String(), err)
	}
	return *output, nil
}

func (r Resource) GetTracks(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	tracksIds model.TracksIds,
) ([]model.Track, error) {
	if err := r.validateTracksIdSize(tracksIds); err != nil {
		return []model.Track{}, err
	}

	url := r.baseUrl + internal.ApiVersion + internal.TracksPath
	queryParams := &model.QueryParams{
		"ids":    tracksIds,
		"market": market,
	}
	output := &model.MultipleTracks{}

	if err := r.httpClient.DoRequest(model.HttpGet, url, queryParams, accessToken, output); err != nil {
		return []model.Track{}, fmt.Errorf("error executing track request for tracks IDs - %s - %w", tracksIds.String(), err)
	}
	return (*output).Tracks, nil
}

func (r Resource) validateTracksIdSize(tracksIds model.TracksIds) error {
	if len(tracksIds) < 1 {
		return fmt.Errorf("error getting track - track id must not be null")
	}
	return nil
}
