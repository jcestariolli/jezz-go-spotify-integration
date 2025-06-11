package albums

import (
	"fmt"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
)

type Service struct {
	authService    *auth.Service
	albumsResource Resource
}

func NewService(
	baseUrl string,
	authService *auth.Service,
) *Service {
	return &Service{
		authService:    authService,
		albumsResource: NewResource(baseUrl),
	}
}

func (s *Service) GetAlbum(
	countryMarketName *string,
	albumId string,
) (model.Album, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return model.Album{}, fmt.Errorf("errror getting album for country %s - invalid country name: %w", *countryMarketName, err)
	}
	getAlbumFn := func() (model.Album, error) {
		return s.albumsResource.GetAlbum(s.authService.GetAppAccessToken(), market, albumId)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getAlbumFn)
}

func (s *Service) GetAlbums(
	countryMarketName *string,
	albumsIds ...string,
) ([]model.Album, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return []model.Album{}, fmt.Errorf("errror getting albums for country %s - invalid country name: %w", *countryMarketName, err)
	}
	getAlbumsFn := func() ([]model.Album, error) {
		return s.albumsResource.GetAlbums(s.authService.GetAppAccessToken(), market, albumsIds...)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getAlbumsFn)
}

func (s *Service) GetAlbumTracks(
	countryMarketName *string,
	limit *model.Limit,
	offset *model.Offset,
	albumId string,
) (model.SimplifiedTracksPaginated, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return model.SimplifiedTracksPaginated{}, fmt.Errorf("errror getting album tracks for country %s - invalid country name: %w", *countryMarketName, err)
	}
	getAlbumTracksFn := func() (model.SimplifiedTracksPaginated, error) {
		return s.albumsResource.GetAlbumTracks(s.authService.GetAppAccessToken(), market, limit, offset, albumId)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getAlbumTracksFn)
}

func (s *Service) GetNewReleases(
	limit *model.Limit,
	offset *model.Offset,
) (model.AlbumsNewRelease, error) {
	getAlbumFn := func() (model.AlbumsNewRelease, error) {
		return s.albumsResource.GetNewReleases(s.authService.GetAppAccessToken(), limit, offset)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getAlbumFn)
}
