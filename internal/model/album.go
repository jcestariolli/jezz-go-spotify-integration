package model

import (
	"strings"

	"github.com/samber/lo"
)

type AlbumsIDs []ID

func (a AlbumsIDs) String() string {
	return strings.Join(lo.Map(a, func(albumID ID, _ int) string {
		return albumID.String()
	}), ",")
}

type AlbumType string

type AlbumGroup string

func (a AlbumGroup) String() string {
	return string(a)
}

type AlbumGroups []AlbumGroup

func (a AlbumGroups) String() string {
	return strings.Join(lo.Map(a, func(albumGroup AlbumGroup, _ int) string {
		return albumGroup.String()
	}), ",")
}

type SimplifiedAlbum struct {
	AlbumType            AlbumType          `json:"album_type"`
	TotalTracks          int                `json:"total_tracks"`
	AvailableMarkets     []AvailableMarket  `json:"available_markets"`
	ExternalURLs         ExternalURLs       `json:"external_ur_ls"`
	Href                 Href               `json:"href"`
	ID                   ID                 `json:"id"`
	Images               []Image            `json:"images"`
	Name                 Name               `json:"name"`
	ReleaseDate          string             `json:"release_date"`
	ReleaseDatePrecision string             `json:"release_date_precision"`
	Restrictions         Restrictions       `json:"restrictions"`
	Type                 Type               `json:"type"`
	URI                  URI                `json:"uri"`
	Artists              []SimplifiedArtist `json:"artists"`
}

type Album struct {
	SimplifiedAlbum
	Tracks      SimplifiedTracksPaginated `json:"tracks"`
	Copyrights  []Copyright               `json:"copyrights"`
	ExternalIDs ExternalIDs               `json:"external_i_ds"`
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

type AlbumsNewRelease struct {
	Albums SimplifiedAlbumsPaginated `json:"albums"`
}
