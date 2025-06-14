package service

import (
	"fmt"
	"jezz-go-spotify-integration/internal/auth"
	"jezz-go-spotify-integration/internal/model"
	"jezz-go-spotify-integration/internal/resource"
	"jezz-go-spotify-integration/internal/utils"

	"github.com/samber/lo"
)

type SpotifyArtistsService struct {
	authService     *auth.Service
	artistsResource resource.ArtistsResource
}

func NewSpotifyArtistsService(
	baseURL string,
	authService *auth.Service,
) ArtistsService {
	return &SpotifyArtistsService{
		authService:     authService,
		artistsResource: resource.NewSpotifyArtistsResource(baseURL),
	}
}

func (s *SpotifyArtistsService) GetArtist(artistID string) (model.Artist, error) {
	getArtistFn := func() (model.Artist, error) {
		return s.artistsResource.GetArtist(s.authService.GetAppAccessToken(), model.ID(artistID))
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistFn)
}

func (s *SpotifyArtistsService) GetArtists(artistIDsStr ...string) ([]model.Artist, error) {
	artistsIDs := lo.Map(artistIDsStr, func(artistID string, _ int) model.ID {
		return model.ID(artistID)
	})
	getArtistsFn := func() ([]model.Artist, error) {
		return s.artistsResource.GetArtists(s.authService.GetAppAccessToken(), artistsIDs)
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistsFn)
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

	getArtistAlbumsFn := func() (model.SimplifiedArtistAlbumsPaginated, error) {
		return s.artistsResource.GetArtistAlbums(s.authService.GetAppAccessToken(), includeGroups, market, _limit, _offset, model.ID(albumID))
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistAlbumsFn)
}

func (s *SpotifyArtistsService) GetArtistTopTracks(
	countryMarketName *string,
	artistID string,
) ([]model.Track, error) {
	market, err := utils.GetMarketByCountryName(countryMarketName)
	if err != nil {
		return []model.Track{}, fmt.Errorf("errror getting artist top-tracks for country %s - invalid country name: %w", *countryMarketName, err)
	}
	getArtistTopTracksFn := func() ([]model.Track, error) {
		return s.artistsResource.GetArtistTopTracks(s.authService.GetAppAccessToken(), market, model.ID(artistID))
	}
	return auth.ExecuteWithAuthRetry(s.authService, getArtistTopTracksFn)

}
