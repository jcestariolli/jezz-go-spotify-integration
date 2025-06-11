package tracks

import (
	"fmt"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/utils"
)

type Service struct {
	authService    *auth.Service
	tracksResource Resource
}

func NewService(
	baseUrl string,
	authService *auth.Service,
) *Service {
	return &Service{
		authService:    authService,
		tracksResource: NewResource(baseUrl),
	}
}

func (c *Service) GetTrack(countryMarketName *string, trackId string) (model.Track, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return model.Track{}, fmt.Errorf("errror getting track for country %s - unknown country! Details: %w", *countryMarketName, err)
	}
	getTrackFn := func() (model.Track, error) {
		return c.tracksResource.GetTrack(c.authService.GetAppAccessToken(), market, trackId)
	}
	return auth.ExecuteWithAuthRetry(c.authService, getTrackFn)
}

func (c *Service) GetTracks(countryMarketName *string, trackIds ...string) ([]model.Track, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return []model.Track{}, fmt.Errorf("errror getting tracks for country %s - unknown country! Details: %w", *countryMarketName, err)
	}
	getTracksFn := func() ([]model.Track, error) {
		return c.tracksResource.GetTracks(c.authService.GetAppAccessToken(), market, trackIds...)
	}
	return auth.ExecuteWithAuthRetry(c.authService, getTracksFn)
}
