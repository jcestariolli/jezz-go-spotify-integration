package artists

import (
	"fmt"
	"github.com/samber/lo"
	"jezz-go-spotify-integration/internal"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
)

type Service struct {
	authService     *auth.Service
	artistsResource internal.ArtistsResource
}

func NewService(
	baseUrl string,
	authService *auth.Service,
) internal.ArtistsService {
	return &Service{
		authService:     authService,
		artistsResource: NewResource(baseUrl),
	}
}

func (s *Service) GetArtist(artistId string) (model.Artist, error) {
	getArtistFn := func() (model.Artist, error) {
		return s.artistsResource.GetArtist(s.authService.GetAppAccessToken(), model.Id(artistId))
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistFn)
}

func (s *Service) GetArtists(artistIdsStr ...string) ([]model.Artist, error) {
	artistsIds := lo.Map(artistIdsStr, func(artistId string, _ int) model.Id {
		return model.Id(artistId)
	})
	getArtistsFn := func() ([]model.Artist, error) {
		return s.artistsResource.GetArtists(s.authService.GetAppAccessToken(), artistsIds)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistsFn)
}

func (s *Service) GetArtistAlbums(
	countryMarketName *string,
	albumTypes *[]string,
	limit *int,
	offset *int,
	albumId string,
) (model.SimplifiedArtistAlbumsPaginated, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return model.SimplifiedArtistAlbumsPaginated{}, fmt.Errorf("errror getting album tracks for country %s - invalid country name: %w", *countryMarketName, err)
	}

	var _limit *model.Limit
	if limit != nil {
		_limit = lo.ToPtr(model.Limit(*limit))
	}
	var _offset *model.Offset
	if offset != nil {
		_offset = lo.ToPtr(model.Offset(*offset))
	}

	var includeGroups *model.AlbumGroups
	if albumTypes != nil && len(*albumTypes) > 0 {
		includeGroups = &model.AlbumGroups{}
		for _, value := range *albumTypes {
			*includeGroups = append(*includeGroups, model.AlbumGroup(value))
		}
	}

	getArtistAlbumsFn := func() (model.SimplifiedArtistAlbumsPaginated, error) {
		return s.artistsResource.GetArtistAlbums(s.authService.GetAppAccessToken(), includeGroups, market, _limit, _offset, model.Id(albumId))
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistAlbumsFn)
}

func (s *Service) GetArtistTopTracks(
	countryMarketName *string,
	artistId string,
) ([]model.Track, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return []model.Track{}, fmt.Errorf("errror getting artist top-tracks for country %s - invalid country name: %w", *countryMarketName, err)
	}
	getArtistTopTracksFn := func() ([]model.Track, error) {
		return s.artistsResource.GetArtistTopTracks(s.authService.GetAppAccessToken(), market, model.Id(artistId))
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistTopTracksFn)

}
