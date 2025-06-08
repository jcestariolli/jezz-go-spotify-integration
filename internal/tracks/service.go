package tracks

import (
	"fmt"
	"github.com/pariz/gountries"
	"github.com/samber/lo"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/model"
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
	market, err := c.getMarketByCountryName(countryMarketName)
	if err != nil {
		return model.Track{}, fmt.Errorf("errror getting track for country %s - unknown country! Details: %w", *countryMarketName, err)
	}
	getTrackFn := func() (model.Track, error) {
		return c.tracksResource.Get(c.authService.GetAppAccessToken(), market, trackId)
	}
	return auth.ExecuteWithAuthRetry(c.authService, getTrackFn)
}

func (c *Service) GetTracks(countryMarketName *string, trackIds ...string) ([]model.Track, error) {
	market, err := c.getMarketByCountryName(countryMarketName)
	if err != nil {
		return []model.Track{}, fmt.Errorf("errror getting tracks for country %s - unknown country! Details: %w", *countryMarketName, err)
	}
	getTracksFn := func() ([]model.Track, error) {
		return c.tracksResource.GetBatch(c.authService.GetAppAccessToken(), market, trackIds...)
	}
	return auth.ExecuteWithAuthRetry(c.authService, getTracksFn)
}

func (c *Service) getMarketByCountryName(countryName *string) (*model.AvailableMarket, error) {
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
