package tracks

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
	tracksResource internal.TracksResource
}

func NewService(
	baseURL string,
	authService *auth.Service,
) internal.TracksService {
	return &Service{
		authService:    authService,
		tracksResource: NewResource(baseURL),
	}
}

func (c *Service) GetTrack(countryMarketName *string, trackID string) (model.Track, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return model.Track{}, fmt.Errorf("errror getting track for country %s - unknown country! Details: %w", *countryMarketName, err)
	}
	getTrackFn := func() (model.Track, error) {
		return c.tracksResource.GetTrack(c.authService.GetAppAccessToken(), market, model.ID(trackID))
	}
	return auth.ExecuteWithAuthRetry(c.authService, getTrackFn)
}

func (c *Service) GetTracks(countryMarketName *string, tracksIDs ...string) ([]model.Track, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return []model.Track{}, fmt.Errorf("errror getting tracks for country %s - unknown country! Details: %w", *countryMarketName, err)
	}

	_tracksIDs := lo.Map(tracksIDs, func(trackID string, _ int) model.ID {
		return model.ID(trackID)
	})

	getTracksFn := func() ([]model.Track, error) {
		return c.tracksResource.GetTracks(c.authService.GetAppAccessToken(), market, _tracksIDs)
	}
	return auth.ExecuteWithAuthRetry(c.authService, getTracksFn)
}
