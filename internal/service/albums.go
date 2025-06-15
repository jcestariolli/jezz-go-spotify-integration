package service

import (
	"fmt"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/resource"
	"jezz-go-spotify-integration/internal/utils"

	"github.com/samber/lo"
)

type SpotifyAlbumsService struct {
	authService    AuthService
	albumsResource resource.AlbumsResource
}

func NewSpotifyAlbumsService(
	baseURL string,
	httpClient client.HTTPClient,
	authService AuthService,
) AlbumsService {
	return &SpotifyAlbumsService{
		authService:    authService,
		albumsResource: resource.NewSpotifyAlbumsResource(httpClient, baseURL),
	}
}

func (s *SpotifyAlbumsService) GetAlbum(
	countryMarketName *string,
	albumID string,
) (model.Album, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return model.Album{}, fmt.Errorf("errror getting album for country %s - invalid country name: %w", *countryMarketName, err)
	}

	result, errA := s.authService.ExecuteWithAuthentication(func(accessToken model.AccessToken) (any, error) {
		return s.albumsResource.GetAlbum(accessToken, market, model.ID(albumID))
	})
	if errA != nil {
		return model.Album{}, errA
	}
	return result.(model.Album), nil
}

func (s *SpotifyAlbumsService) GetAlbums(
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
	result, errA := s.authService.ExecuteWithAuthentication(func(accessToken model.AccessToken) (any, error) {
		return s.albumsResource.GetAlbums(accessToken, market, _albumsIDs)
	})
	if errA != nil {
		return []model.Album{}, errA
	}
	return result.([]model.Album), nil
}

func (s *SpotifyAlbumsService) GetAlbumTracks(
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

	result, errA := s.authService.ExecuteWithAuthentication(func(accessToken model.AccessToken) (any, error) {
		return s.albumsResource.GetAlbumTracks(accessToken, market, _limit, _offset, model.ID(albumID))
	})
	if errA != nil {
		return model.SimplifiedTracksPaginated{}, errA
	}
	return result.(model.SimplifiedTracksPaginated), nil

}

func (s *SpotifyAlbumsService) GetNewReleases(
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

	result, errA := s.authService.ExecuteWithAuthentication(func(accessToken model.AccessToken) (any, error) {
		return s.albumsResource.GetNewReleases(accessToken, _limit, _offset)
	})
	if errA != nil {
		return model.AlbumsNewRelease{}, errA
	}
	return result.(model.AlbumsNewRelease), nil
}
