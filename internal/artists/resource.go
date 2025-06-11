package artists

import (
	"fmt"
	"github.com/samber/lo"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
	"net/http"
	"strings"
)

const (
	apiVersion      = "/v1"
	artistsResource = "/artists"
	albumsResource  = "/albums"
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

func (r Resource) GetArtist(
	accessToken model.AccessToken,
	artistId string,
) (model.Artist, error) {
	req, cErr := utils.CreateHttpRequest(utils.HttpGet, r.baseUrl+apiVersion, artistsResource+"/"+artistId, map[string]string{}, accessToken)
	if cErr != nil {
		return model.Artist{}, fmt.Errorf("error creating artist request for astist ID - %s - %w", artistId, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return model.Artist{}, fmt.Errorf("error connecting to artist client for astist ID - %s - %w", artistId, reqErr)
	}

	if vErr := utils.ValidateHttpResponseStatus(resp); vErr != nil {
		return model.Artist{}, vErr
	}
	output := &model.Artist{}
	if pErr := utils.ParseHttpResponse(resp, output); pErr != nil {
		return model.Artist{}, fmt.Errorf("error parsing response from resource for astist ID - %s - %w", artistId, pErr)
	}
	return *output, nil
}

func (r Resource) GetArtists(
	accessToken model.AccessToken,
	artistsIds ...string,
) ([]model.Artist, error) {
	if err := r.validateArtistsIdSize(artistsIds); err != nil {
		return []model.Artist{}, err
	}
	artistsIdsStr := strings.Join(artistsIds, ",")
	queryParameters := map[string]string{
		"ids": artistsIdsStr,
	}

	req, cErr := utils.CreateHttpRequest(utils.HttpGet, r.baseUrl+apiVersion, artistsResource, queryParameters, accessToken)
	if cErr != nil {
		return []model.Artist{}, fmt.Errorf("error creating artist request for astists IDs - %s - %w", artistsIdsStr, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return []model.Artist{}, fmt.Errorf("error connecting to artist client for astists IDs - %s - %w", artistsIdsStr, reqErr)
	}

	if vErr := utils.ValidateHttpResponseStatus(resp); vErr != nil {
		return []model.Artist{}, vErr
	}

	output := &model.MultipleArtists{}
	if pErr := utils.ParseHttpResponse(resp, output); pErr != nil {
		return []model.Artist{}, fmt.Errorf("error parsing response from resource for astists ID - %s - %w", artistsIdsStr, pErr)
	}
	return output.Artists, nil
}

func (r Resource) GetArtistAlbums(
	accessToken model.AccessToken,
	includeGroups []model.AlbumGroup,
	market *model.AvailableMarket,
	limit *model.Limit,
	offset *model.Offset,
	artistId string,
) (model.SimplifiedArtistAlbumsPaginated, error) {
	queryParameters := map[string]string{}
	if len(includeGroups) > 0 {
		queryParameters["include_groups"] = strings.Join(
			lo.Map(
				includeGroups,
				func(albumGroup model.AlbumGroup, _ int) string { return albumGroup.String() },
			),
			",",
		)
	}
	params := []model.Pair[string, model.StringEvaluator]{
		{"market", market},
		{"limit", limit},
		{"market", offset},
	}
	queryParameters = utils.AppendQueryParams(queryParameters, params...)

	req, cErr := utils.CreateHttpRequest(utils.HttpGet, r.baseUrl+apiVersion, artistsResource+"/"+artistId+albumsResource, queryParameters, accessToken)
	if cErr != nil {
		return model.SimplifiedArtistAlbumsPaginated{}, fmt.Errorf("error creating artist albums request for astist ID - %s - %w", artistId, cErr)
	}

	resp, reqErr := (&http.Client{}).Do(req)
	if reqErr != nil {
		return model.SimplifiedArtistAlbumsPaginated{}, fmt.Errorf("error connecting to artist albums client for astist ID - %s - %w", artistId, reqErr)
	}

	if vErr := utils.ValidateHttpResponseStatus(resp); vErr != nil {
		return model.SimplifiedArtistAlbumsPaginated{}, vErr
	}
	output := &model.SimplifiedArtistAlbumsPaginated{}
	if pErr := utils.ParseHttpResponse(resp, output); pErr != nil {
		return model.SimplifiedArtistAlbumsPaginated{}, fmt.Errorf("error parsing response from resource for astist albums ID - %s - %w", artistId, pErr)
	}
	return *output, nil
}

func (r Resource) validateArtistsIdSize(artistIds []string) error {
	if len(artistIds) < 1 {
		return fmt.Errorf("error getting artist - artist id must not be null")
	}
	return nil
}
