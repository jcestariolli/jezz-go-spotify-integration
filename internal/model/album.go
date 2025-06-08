package model

type AlbumType string

const (
	DefaultAlbumType     AlbumType = "album"
	SingleAlbumType      AlbumType = "single"
	AppearsOnAlgumType   AlbumType = "appears_on"
	CompilationAlbumType AlbumType = "compilation"
)

type Album struct {
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
	Tracks               Tracks             `json:"tracks"`
	Copyrights           []Copyright        `json:"copyrights"`
	ExternalIds          ExternalIds        `json:"external_ids"`
	Label                string             `json:"label"`
	Popularity           int                `json:"popularity"`
}

type Albums []Album

type MultipleAlbums struct {
	Albums Albums `json:"albums"`
}
