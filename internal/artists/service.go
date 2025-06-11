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

func (s *Service) GetArtist(artistId string) (model.Artist, error) {
	getArtistFn := func() (model.Artist, error) {
		return s.artistsResource.GetArtist(s.authService.GetAppAccessToken(), artistId)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistFn)
}

func (s *Service) GetArtists(artistIds ...string) ([]model.Artist, error) {
	getArtistsFn := func() ([]model.Artist, error) {
		return s.artistsResource.GetArtists(s.authService.GetAppAccessToken(), artistIds...)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistsFn)
}
