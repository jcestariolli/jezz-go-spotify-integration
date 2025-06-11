package artists

import (
	"fmt"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
)

type Service struct {
	authService     *auth.Service
	artistsResource Resource
}

func NewService(
	baseUrl string,
	authService *auth.Service,
) *Service {
	return &Service{
		authService:     authService,
		artistsResource: NewResource(baseUrl),
	}
}

func (s *Service) GetArtist(artistId string) (model.Artist, error) {
	getArtistFn := func() (model.Artist, error) {
		return s.artistsResource.GetArtist(s.authService.GetAppAccessToken(), artistId)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistFn)
}

func (s *Service) GetArtists(artistIds ...string) ([]model.Artist, error) {
	getArtistsFn := func() ([]model.Artist, error) {
		return s.artistsResource.GetArtists(s.authService.GetAppAccessToken(), artistIds...)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistsFn)
}

func (s *Service) GetArtistAlbums(
	countryMarketName *string,
	includeGroups []model.AlbumGroup,
	limit *model.Limit,
	offset *model.Offset,
	albumId string,
) (model.SimplifiedArtistAlbumsPaginated, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return model.SimplifiedArtistAlbumsPaginated{}, fmt.Errorf("errror getting album tracks for country %s - invalid country name: %w", *countryMarketName, err)
	}
	getArtistAlbumsFn := func() (model.SimplifiedArtistAlbumsPaginated, error) {
		return s.artistsResource.GetArtistAlbums(s.authService.GetAppAccessToken(), includeGroups, market, limit, offset, albumId)
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
		return s.artistsResource.GetArtistTopTracks(s.authService.GetAppAccessToken(), market, artistId)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistTopTracksFn)

}
