package artists

import (
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/model"
)

type Service struct {
	authService     *auth.Service
	artistsResource Resource
}

func NewService(
	baseUrl string,
	authService *auth.Service,
) *Service {
	return &Service{
		authService:     authService,
		artistsResource: NewResource(baseUrl),
	}
}

func (c *Service) GetArtist(artistId string) (model.Artist, error) {
	getArtistFn := func() (model.Artist, error) {
		return c.artistsResource.Get(c.authService.GetAppAccessToken(), artistId)
	}
	return auth.ExecuteWithAuthRetry(c.authService, getArtistFn)
}

func (c *Service) GetArtists(artistIds ...string) (model.Artists, error) {
	getArtistsFn := func() (model.Artists, error) {
		return c.artistsResource.GetBatch(c.authService.GetAppAccessToken(), artistIds...)
	}
	return auth.ExecuteWithAuthRetry(c.authService, getArtistsFn)
}
