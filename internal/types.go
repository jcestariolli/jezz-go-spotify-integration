package internal

import "jezz-go-spotify-integration/internal/model"

type AlbumsResource interface {
	GetAlbum(accessToken model.AccessToken, market *model.AvailableMarket, albumId model.Id) (model.Album, error)
	GetAlbums(accessToken model.AccessToken, market *model.AvailableMarket, albumsIds model.AlbumsIds) ([]model.Album, error)
	GetAlbumTracks(accessToken model.AccessToken, market *model.AvailableMarket, limit *model.Limit, offset *model.Offset, albumId model.Id) (model.SimplifiedTracksPaginated, error)
	GetNewReleases(accessToken model.AccessToken, limit *model.Limit, offset *model.Offset) (model.AlbumsNewRelease, error)
}

type AlbumsService interface {
	GetAlbum(countryMarketName *string, albumId string) (model.Album, error)
	GetAlbums(countryMarketName *string, albumsIds ...string) ([]model.Album, error)
	GetAlbumTracks(countryMarketName *string, limit *int, offset *int, albumId string) (model.SimplifiedTracksPaginated, error)
	GetNewReleases(limit *int, offset *int) (model.AlbumsNewRelease, error)
}

type ArtistsResource interface {
	GetArtist(accessToken model.AccessToken, artistId model.Id) (model.Artist, error)
	GetArtists(accessToken model.AccessToken, artistsIds model.ArtistsIds) ([]model.Artist, error)
	GetArtistAlbums(accessToken model.AccessToken, includeGroups *model.AlbumGroups, market *model.AvailableMarket, limit *model.Limit, offset *model.Offset, artistId model.Id) (model.SimplifiedArtistAlbumsPaginated, error)
	GetArtistTopTracks(accessToken model.AccessToken, market *model.AvailableMarket, artistId model.Id) ([]model.Track, error)
}

type ArtistsService interface {
	GetArtist(artistId string) (model.Artist, error)
	GetArtists(artistIdsStr ...string) ([]model.Artist, error)
	GetArtistAlbums(countryMarketName *string, albumTypes *[]string, limit *int, offset *int, albumId string) (model.SimplifiedArtistAlbumsPaginated, error)
	GetArtistTopTracks(countryMarketName *string, artistId string) ([]model.Track, error)
}

type TracksResource interface {
	GetTrack(accessToken model.AccessToken, market *model.AvailableMarket, trackId model.Id) (model.Track, error)
	GetTracks(accessToken model.AccessToken, market *model.AvailableMarket, tracksIds model.TracksIds) ([]model.Track, error)
}

type TracksService interface {
	GetTrack(countryMarketName *string, trackId string) (model.Track, error)
	GetTracks(countryMarketName *string, tracksIds ...string) ([]model.Track, error)
}
