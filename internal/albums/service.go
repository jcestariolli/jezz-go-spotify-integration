package albums

import (
	"fmt"
	"github.com/pariz/gountries"
	"github.com/samber/lo"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/model"
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
	market, err := s.getMarketByCountryName(countryMarketName)
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
	market, err := s.getMarketByCountryName(countryMarketName)
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
	market, err := s.getMarketByCountryName(countryMarketName)
	if err != nil {
		return model.SimplifiedTracksPaginated{}, fmt.Errorf("errror getting album tracks for country %s - invalid country name: %w", *countryMarketName, err)
	}
	getAlbumFn := func() (model.SimplifiedTracksPaginated, error) {
		return s.albumsResource.GetAlbumTracks(s.authService.GetAppAccessToken(), market, limit, offset, albumId)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getAlbumFn)
}

func (s *Service) getMarketByCountryName(countryName *string) (*model.AvailableMarket, error) {
	var market *model.AvailableMarket
	if countryName != nil {
		country, err := gountries.New().FindCountryByName(*countryName)
		if err != nil {
			return nil, err
		}
		market = lo.ToPtr(model.AvailableMarket(country.Alpha2))

	}
	return market, nil
}
