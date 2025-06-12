package model

import (
	"github.com/samber/lo"
	"strings"
)

type ArtistsIds []Id

func (a ArtistsIds) String() string {
	return strings.Join(lo.Map(a, func(artistId Id, _ int) string {
		return artistId.String()
	}), ",")
}

type SimplifiedArtist struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         Href         `json:"href"`
	Id           Id           `json:"id"`
	Name         Name         `json:"name"`
	Type         Type         `json:"type"`
	Uri          Uri          `json:"uri"`
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
