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

func (c *Service) GetAlbum(countryMarketName *string, albumId string) (model.Album, error) {
	market, err := c.getMarketByCountryName(countryMarketName)
	if err != nil {
		return model.Album{}, fmt.Errorf("errror getting album for country %s - unknown country! Details: %w", *countryMarketName, err)
	}
	getAlbumFn := func() (model.Album, error) {
		return c.albumsResource.Get(c.authService.GetAppAccessToken(), market, albumId)
	}
	return auth.ExecuteWithAuthRetry(c.authService, getAlbumFn)
}

func (c *Service) GetAlbums(countryMarketName *string, albumIds ...string) ([]model.Album, error) {
	market, err := c.getMarketByCountryName(countryMarketName)
	if err != nil {
		return []model.Album{}, fmt.Errorf("errror getting albums for country %s - unknown country! Details: %w", *countryMarketName, err)
	}
	getAlbumsFn := func() ([]model.Album, error) {
		return c.albumsResource.GetBatch(c.authService.GetAppAccessToken(), market, albumIds...)
	}
	return auth.ExecuteWithAuthRetry(c.authService, getAlbumsFn)
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
