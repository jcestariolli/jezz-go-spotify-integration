package model

type AlbumType string

type AlbumGroup string

const (
	DefaultAlbumGroup     AlbumGroup = "album"
	SingleAlbumGroup      AlbumGroup = "single"
	AppearsOnAlgumGroup   AlbumGroup = "appears_on"
	CompilationAlbumGroup AlbumGroup = "compilation"
)

type SimplifiedAlbum struct {
	AlbumType            AlbumType          `json:"album_type"`
	TotalTracks          int                `json:"total_tracks"`
	AvailableMarkets     []AvailableMarket  `json:"available_markets"`
	ExternalUrls         ExternalUrls       `json:"external_urls"`
	Href                 Href               `json:"href"`
	Id                   Id                 `json:"id"`
	Images               []Image            `json:"images"`
	Name                 Name               `json:"name"`
	ReleaseDate          string             `json:"release_date"`
	ReleaseDatePrecision string             `json:"release_date_precision"`
	Restrictions         Restrictions       `json:"restrictions"`
	Type                 Type               `json:"type"`
	Uri                  Uri                `json:"uri"`
	Artists              []SimplifiedArtist `json:"artists"`
}

type SimplifiedArtistAlbum struct {
	SimplifiedAlbum
	AlbumGroup AlbumGroup `json:"album_group"`
}

type Album struct {
	SimplifiedAlbum
	Tracks      SimplifiedTracksPaginated `json:"tracks"`
	Copyrights  []Copyright               `json:"copyrights"`
	ExternalIds ExternalIds               `json:"external_ids"`
	Label       string                    `json:"label"`
	Popularity  int                       `json:"popularity"`
}

type MultipleAlbums struct {
	Albums []Album `json:"albums"`
}

type SimplifiedAlbumsPaginated struct {
	Pagination
	Items []SimplifiedAlbum `json:"items"`
}

type SimplifiedArtistAlbumsPaginated struct {
	Pagination
	Items []SimplifiedArtistAlbum `json:"items"`
}

type AlbumsNewRelease struct {
	Albums SimplifiedAlbumsPaginated `json:"albums"`
}
