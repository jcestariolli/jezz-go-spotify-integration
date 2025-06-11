package tracks

import (
	"fmt"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
	"strings"
)

const (
	apiVersion     = "/v1"
	tracksResource = "/tracks"
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

func (r Resource) GetTrack(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	trackId string,
) (model.Track, error) {
	url := r.baseUrl + apiVersion + tracksResource + "/" + trackId
	queryParams := map[string]string{}
	params := []model.Pair[string, model.StringEvaluator]{
		{"market", market},
	}
	queryParams = utils.AppendQueryParams(queryParams, params...)
	output := &model.Track{}

	if err := utils.DoGetRequestAndValidateSuccess(url, queryParams, accessToken, output); err != nil {
		return model.Track{}, fmt.Errorf("error executing track request for track ID - %s - %w", trackId, err)
	}
	return *output, nil
}

func (r Resource) GetTracks(
	accessToken model.AccessToken,
	market *model.AvailableMarket,
	tracksIds ...string,
) ([]model.Track, error) {
	if err := r.validateTracksIdSize(tracksIds); err != nil {
		return []model.Track{}, err
	}

	url := r.baseUrl + apiVersion + tracksResource
	tracksIdsStr := strings.Join(tracksIds, ",")
	queryParams := map[string]string{
		"ids": tracksIdsStr,
	}
	params := []model.Pair[string, model.StringEvaluator]{
		{"market", market},
	}
	queryParams = utils.AppendQueryParams(queryParams, params...)
	output := &model.MultipleTracks{}

	if err := utils.DoGetRequestAndValidateSuccess(url, queryParams, accessToken, output); err != nil {
		return []model.Track{}, fmt.Errorf("error executingtrack request for tracks IDs - %s - %w", tracksIdsStr, err)
	}
	return (*output).Tracks, nil
}

func (r Resource) validateTracksIdSize(trackIds []string) error {
	if len(trackIds) < 1 {
		return fmt.Errorf("error getting track - track id must not be null")
	}
	return nil
}
