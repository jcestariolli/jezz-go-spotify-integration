package service

import "jezz-go-spotify-integration/internal/model"

type FnWithAuthentication func(accessToken model.AccessToken) (any, error)

type AuthService interface {
	ExecuteWithAuthentication(fn FnWithAuthentication) (any, error)
}

type AlbumsService interface {
	GetAlbum(countryMarketName *string, albumID string) (model.Album, error)
	GetAlbums(countryMarketName *string, albumsIDs ...string) ([]model.Album, error)
	GetAlbumTracks(countryMarketName *string, limit *int, offset *int, albumID string) (model.SimplifiedTracksPaginated, error)
	GetNewReleases(limit *int, offset *int) (model.AlbumsNewRelease, error)
}

type ArtistsService interface {
	GetArtist(artistID string) (model.Artist, error)
	GetArtists(artistIDsStr ...string) ([]model.Artist, error)
	GetArtistAlbums(countryMarketName *string, albumTypes *[]string, limit *int, offset *int, albumID string) (model.SimplifiedArtistAlbumsPaginated, error)
	GetArtistTopTracks(countryMarketName *string, artistID string) ([]model.Track, error)
}

type TracksService interface {
	GetTrack(countryMarketName *string, trackID string) (model.Track, error)
	GetTracks(countryMarketName *string, tracksIDs ...string) ([]model.Track, error)
}
