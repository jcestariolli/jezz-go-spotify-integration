package service

import (
	"fmt"
	"jezz-go-spotify-integration/internal/client"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/resource"
	"jezz-go-spotify-integration/internal/utils"

	"github.com/samber/lo"
)

type SpotifyTracksService struct {
	authService    AuthService
	tracksResource resource.TracksResource
}

func NewSpotifyTracksService(
	baseURL string,
	httpAPIClient client.HTTPApiClient,
	authService AuthService,
) TracksService {
	return &SpotifyTracksService{
		authService:    authService,
		tracksResource: resource.NewSpotifyTracksResource(httpAPIClient, baseURL),
	}
}

func (s *SpotifyTracksService) GetTrack(countryMarketName *string, trackID string) (model.Track, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return model.Track{}, fmt.Errorf("errror getting track for country %s - unknown country! Details: %w", *countryMarketName, err)
	}

	result, errA := s.authService.ExecuteWithAuthentication(func(accessToken model.AccessToken) (any, error) {
		return s.tracksResource.GetTrack(accessToken, market, model.ID(trackID))
	})
	if errA != nil {
		return model.Track{}, errA
	}
	return result.(model.Track), nil
}

func (s *SpotifyTracksService) GetTracks(countryMarketName *string, tracksIDs ...string) ([]model.Track, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return []model.Track{}, fmt.Errorf("errror getting tracks for country %s - unknown country! Details: %w", *countryMarketName, err)
	}

	_tracksIDs := lo.Map(tracksIDs, func(trackID string, _ int) model.ID {
		return model.ID(trackID)
	})

	result, errA := s.authService.ExecuteWithAuthentication(func(accessToken model.AccessToken) (any, error) {
		return s.tracksResource.GetTracks(accessToken, market, _tracksIDs)
	})
	if errA != nil {
		return []model.Track{}, errA
	}
	return result.([]model.Track), nil
}
