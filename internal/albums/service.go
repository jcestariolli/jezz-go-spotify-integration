package albums

import (
	"fmt"
	"jezz-go-spotify-integration/internal"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"

	"github.com/samber/lo"
)

type Service struct {
	authService    *auth.Service
	albumsResource internal.AlbumsResource
}

func NewService(
	baseURL string,
	authService *auth.Service,
) internal.AlbumsService {
	return &Service{
		authService:    authService,
		albumsResource: NewResource(baseURL),
	}
}

func (s *Service) GetAlbum(
	countryMarketName *string,
	albumID string,
) (model.Album, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return model.Album{}, fmt.Errorf("errror getting album for country %s - invalid country name: %w", *countryMarketName, err)
	}
	getAlbumFn := func() (model.Album, error) {
		return s.albumsResource.GetAlbum(s.authService.GetAppAccessToken(), market, model.ID(albumID))
	}
	return auth.ExecuteWithAuthRetry(s.authService, getAlbumFn)
}

func (s *Service) GetAlbums(
	countryMarketName *string,
	albumsIDs ...string,
) ([]model.Album, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return []model.Album{}, fmt.Errorf("errror getting albums for country %s - invalid country name: %w", *countryMarketName, err)
	}

	_albumsIDs := lo.Map(albumsIDs, func(albumID string, _ int) model.ID {
		return model.ID(albumID)
	})
	getAlbumsFn := func() ([]model.Album, error) {
		return s.albumsResource.GetAlbums(s.authService.GetAppAccessToken(), market, _albumsIDs)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getAlbumsFn)
}

func (s *Service) GetAlbumTracks(
	countryMarketName *string,
	limit *int,
	offset *int,
	albumID string,
) (model.SimplifiedTracksPaginated, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return model.SimplifiedTracksPaginated{}, fmt.Errorf("errror getting album tracks for country %s - invalid country name: %w", *countryMarketName, err)
	}

	var _limit *model.Limit
	if limit != nil {
		_limit = lo.ToPtr(model.Limit(*limit))
	}
	var _offset *model.Offset
	if offset != nil {
		_offset = lo.ToPtr(model.Offset(*offset))
	}

	getAlbumTracksFn := func() (model.SimplifiedTracksPaginated, error) {
		return s.albumsResource.GetAlbumTracks(s.authService.GetAppAccessToken(), market, _limit, _offset, model.ID(albumID))
	}
	return auth.ExecuteWithAuthRetry(s.authService, getAlbumTracksFn)
}

func (s *Service) GetNewReleases(
	limit *int,
	offset *int,
) (model.AlbumsNewRelease, error) {
	var _limit *model.Limit
	if limit != nil {
		_limit = lo.ToPtr(model.Limit(*limit))
	}
	var _offset *model.Offset
	if offset != nil {
		_offset = lo.ToPtr(model.Offset(*offset))
	}

	getAlbumFn := func() (model.AlbumsNewRelease, error) {
		return s.albumsResource.GetNewReleases(s.authService.GetAppAccessToken(), _limit, _offset)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getAlbumFn)
}
