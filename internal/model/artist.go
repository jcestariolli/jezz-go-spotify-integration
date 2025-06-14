package model

import (
	"strings"

	"github.com/samber/lo"
)

type ArtistsIDs []ID

func (a ArtistsIDs) String() string {
	return strings.Join(lo.Map(a, func(artistID ID, _ int) string {
		return artistID.String()
	}), ",")
}

type SimplifiedArtist struct {
	ExternalURLs ExternalURLs `json:"external_ur_ls"`
	Href         Href         `json:"href"`
	ID           ID           `json:"id"`
	Name         Name         `json:"name"`
	Type         Type         `json:"type"`
	URI          URI          `json:"uri"`
}

type Artist struct {
	SimplifiedArtist
	Followers  Followers `json:"followers"`
	Genres     Genres    `json:"genres"`
	Images     []Image   `json:"images"`
	Popularity int       `json:"popularity"`
}

type MultipleArtists struct {
	Artists []Artist `json:"artists"`
}

type SimplifiedArtistAlbum struct {
	SimplifiedAlbum
	AlbumGroup AlbumGroup `json:"album_group"`
}

type SimplifiedArtistAlbumsPaginated struct {
	Pagination
	Items []SimplifiedArtistAlbum `json:"items"`
}
