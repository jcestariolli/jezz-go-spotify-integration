package artists

import (
	"fmt"
	"github.com/samber/lo"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
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
	url := r.baseUrl + apiVersion + artistsResource + "/" + artistId
	queryParams := map[string]string{}
	output := &model.Artist{}

	if err := utils.DoGetRequestAndValidateSuccess(url, queryParams, accessToken, output); err != nil {
		return model.Artist{}, fmt.Errorf("error executing artist request for astist ID - %s - %w", artistId, err)
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

	url := r.baseUrl + apiVersion + artistsResource
	queryParams := map[string]string{
		"ids": artistsIdsStr,
	}
	output := &model.MultipleArtists{}

	if err := utils.DoGetRequestAndValidateSuccess(url, queryParams, accessToken, output); err != nil {
		return []model.Artist{}, fmt.Errorf("error executing artist request for astists IDs - %s - %w", artistsIdsStr, err)
	}
	return (*output).Artists, nil
}

func (r Resource) GetArtistAlbums(
	accessToken model.AccessToken,
	includeGroups []model.AlbumGroup,
	market *model.AvailableMarket,
	limit *model.Limit,
	offset *model.Offset,
	artistId string,
) (model.SimplifiedArtistAlbumsPaginated, error) {
	url := r.baseUrl + apiVersion + artistsResource + "/" + artistId + albumsResource
	queryParams := map[string]string{}
	if len(includeGroups) > 0 {
		queryParams["include_groups"] = strings.Join(
			lo.Map(includeGroups, func(albumGroup model.AlbumGroup, _ int) string { return albumGroup.String() }),
			",",
		)
	}
	params := []model.Pair[string, model.StringEvaluator]{
		{"market", market},
		{"limit", limit},
		{"market", offset},
	}
	queryParams = utils.AppendQueryParams(queryParams, params...)
	output := &model.SimplifiedArtistAlbumsPaginated{}

	if err := utils.DoGetRequestAndValidateSuccess(url, queryParams, accessToken, output); err != nil {
		return model.SimplifiedArtistAlbumsPaginated{}, fmt.Errorf("error executing artist albums request for astist ID - %s - %w", artistId, err)
	}
	return *output, nil
}

func (r Resource) validateArtistsIdSize(artistIds []string) error {
	if len(artistIds) < 1 {
		return fmt.Errorf("error getting artist - artist id must not be null")
	}
	return nil
}
