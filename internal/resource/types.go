package resource

import "jezz-go-spotify-integration/internal/model"

type AlbumsResource interface {
	GetAlbum(accessToken model.AccessToken, market *model.AvailableMarket, albumID model.ID) (model.Album, error)
	GetAlbums(accessToken model.AccessToken, market *model.AvailableMarket, albumsIDs model.AlbumsIDs) ([]model.Album, error)
	GetAlbumTracks(accessToken model.AccessToken, market *model.AvailableMarket, limit *model.Limit, offset *model.Offset, albumID model.ID) (model.SimplifiedTracksPaginated, error)
	GetNewReleases(accessToken model.AccessToken, limit *model.Limit, offset *model.Offset) (model.AlbumsNewRelease, error)
}

type ArtistsResource interface {
	GetArtist(accessToken model.AccessToken, artistID model.ID) (model.Artist, error)
	GetArtists(accessToken model.AccessToken, artistsIDs model.ArtistsIDs) ([]model.Artist, error)
	GetArtistAlbums(accessToken model.AccessToken, includeGroups *model.AlbumGroups, market *model.AvailableMarket, limit *model.Limit, offset *model.Offset, artistID model.ID) (model.SimplifiedArtistAlbumsPaginated, error)
	GetArtistTopTracks(accessToken model.AccessToken, market *model.AvailableMarket, artistID model.ID) ([]model.Track, error)
}

type TracksResource interface {
	GetTrack(accessToken model.AccessToken, market *model.AvailableMarket, trackID model.ID) (model.Track, error)
	GetTracks(accessToken model.AccessToken, market *model.AvailableMarket, tracksIDs model.TracksIDs) ([]model.Track, error)
}
