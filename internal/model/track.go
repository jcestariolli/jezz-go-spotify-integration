package model

import (
	"github.com/samber/lo"
	"strings"
)

type TracksIds []Id

func (a TracksIds) String() string {
	return strings.Join(lo.Map(a, func(trackId Id, _ int) string {
		return trackId.String()
	}), ",")
}

type SimplifiedTrack struct {
	Artists          []SimplifiedArtist `json:"artists"`
	AvailableMarkets []AvailableMarket  `json:"available_markets"`
	DiscNumber       int                `json:"disc_number"`
	DurationMs       int                `json:"duration_ms"`
	Explicit         bool               `json:"explicit"`
	ExternalUrls     ExternalUrls       `json:"external_urls"`
	Href             Href               `json:"href"`
	Id               Id                 `json:"id"`
	IsPlayable       bool               `json:"is_playable"`
	LinkedFrom       LinkedFrom         `json:"linked_from"`
	Restrictions     Restrictions       `json:"restrictions"`
	Name             Name               `json:"name"`
	TrackNumber      int                `json:"track_number"`
	Type             Type               `json:"type"`
	Uri              Uri                `json:"uri"`
	IsLocal          bool               `json:"is_local"`
}

type Track struct {
	SimplifiedTrack
	Album       SimplifiedAlbum `json:"album"`
	ExternalIds ExternalIds     `json:"external_ids"`
	Popularity  int             `json:"popularity"`
}

type MultipleTracks struct {
	Tracks []Track `json:"tracks"`
}

type SimplifiedTracksPaginated struct {
	Pagination
	Items []SimplifiedTrack `json:"items"`
}
