package tracks

import (
	"fmt"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
	"net/http"
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
	queryParameters := map[string]string{}
	params := []model.Pair[string, model.StringEvaluator]{
		{"market", market},
	}
	queryParameters = utils.AppendQueryParams(queryParameters, params...)

	req, cErr := utils.CreateHttpRequest(utils.HttpGet, r.baseUrl+apiVersion+tracksResource, "/"+trackId, queryParameters, accessToken)
	if cErr != nil {
		return model.Track{}, fmt.Errorf("error creating track request for track ID - %s - %w", trackId, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return model.Track{}, fmt.Errorf("error connecting to track client for track ID - %s - %w", trackId, reqErr)
	}

	if vErr := utils.ValidateHttpResponseStatus(resp); vErr != nil {
		return model.Track{}, vErr
	}
	output := &model.Track{}
	if pErr := utils.ParseHttpResponse(resp, output); pErr != nil {
		return model.Track{}, fmt.Errorf("error parsing response from resource for track ID - %s - %w", trackId, pErr)
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

	tracksIdsStr := strings.Join(tracksIds, ",")
	queryParameters := map[string]string{
		"ids": tracksIdsStr,
	}
	params := []model.Pair[string, model.StringEvaluator]{
		{"market", market},
	}
	queryParameters = utils.AppendQueryParams(queryParameters, params...)

	req, cErr := utils.CreateHttpRequest(utils.HttpGet, r.baseUrl+apiVersion+tracksResource, "", queryParameters, accessToken)
	if cErr != nil {
		return []model.Track{}, fmt.Errorf("error creating track request for tracks IDs - %s - %w", tracksIdsStr, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return []model.Track{}, fmt.Errorf("error connecting to track client for tracks IDs - %s - %w", tracksIdsStr, reqErr)
	}

	if vErr := utils.ValidateHttpResponseStatus(resp); vErr != nil {
		return []model.Track{}, vErr
	}

	output := &model.MultipleTracks{}
	if pErr := utils.ParseHttpResponse(resp, output); pErr != nil {
		return []model.Track{}, fmt.Errorf("error parsing response from resource for tracks ID - %s - %w", tracksIdsStr, pErr)
	}
	return output.Tracks, nil

}

func (r Resource) validateTracksIdSize(trackIds []string) error {
	if len(trackIds) < 1 {
		return fmt.Errorf("error getting track - track id must not be null")
	}
	return nil
}
