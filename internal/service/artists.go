package service

import (
	"fmt"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/resource"
	"jezz-go-spotify-integration/internal/utils"

	"github.com/samber/lo"
)

type SpotifyArtistsService struct {
	authService     AuthService
	artistsResource resource.ArtistsResource
}

func NewSpotifyArtistsService(
	baseURL string,
	authService AuthService,
) ArtistsService {
	return &SpotifyArtistsService{
		authService:     authService,
		artistsResource: resource.NewSpotifyArtistsResource(baseURL),
	}
}

func (s *SpotifyArtistsService) GetArtist(artistID string) (model.Artist, error) {
	result, errA := s.authService.ExecuteWithAuthentication(func(accessToken model.AccessToken) (any, error) {
		return s.artistsResource.GetArtist(accessToken, model.ID(artistID))
	})
	if errA != nil {
		return model.Artist{}, errA
	}
	return result.(model.Artist), nil
}

func (s *SpotifyArtistsService) GetArtists(artistIDsStr ...string) ([]model.Artist, error) {
	artistsIDs := lo.Map(artistIDsStr, func(artistID string, _ int) model.ID {
		return model.ID(artistID)
	})
	result, errA := s.authService.ExecuteWithAuthentication(func(accessToken model.AccessToken) (any, error) {
		return s.artistsResource.GetArtists(accessToken, artistsIDs)
	})
	if errA != nil {
		return []model.Artist{}, errA
	}
	return result.([]model.Artist), nil
}

func (s *SpotifyArtistsService) GetArtistAlbums(
	countryMarketName *string,
	albumTypes *[]string,
	limit *int,
	offset *int,
	albumID string,
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

	result, errA := s.authService.ExecuteWithAuthentication(func(accessToken model.AccessToken) (any, error) {
		return s.artistsResource.GetArtistAlbums(accessToken, includeGroups, market, _limit, _offset, model.ID(albumID))
	})
	if errA != nil {
		return model.SimplifiedArtistAlbumsPaginated{}, errA
	}
	return result.(model.SimplifiedArtistAlbumsPaginated), nil
}

func (s *SpotifyArtistsService) GetArtistTopTracks(
	countryMarketName *string,
	artistID string,
) ([]model.Track, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return []model.Track{}, fmt.Errorf("errror getting artist top-tracks for country %s - invalid country name: %w", *countryMarketName, err)
	}

	result, errA := s.authService.ExecuteWithAuthentication(func(accessToken model.AccessToken) (any, error) {
		return s.artistsResource.GetArtistTopTracks(accessToken, market, model.ID(artistID))
	})
	if errA != nil {
		return []model.Track{}, errA
	}
	return result.([]model.Track), nil
}
